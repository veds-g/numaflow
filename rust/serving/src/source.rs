use std::collections::HashMap;
use std::sync::Arc;
use std::time::Duration;

use async_nats::jetstream::Context;
use bytes::Bytes;
use tokio::sync::{mpsc, oneshot};
use tokio::time::Instant;

use crate::app::orchestrator::OrchestratorState as CallbackState;
use crate::app::store::cbstore::jetstreamstore::JetStreamCallbackStore;
use crate::app::store::datastore::jetstream::JetStreamDataStore;
use crate::app::store::datastore::user_defined::UserDefinedStore;
use crate::app::tracker::MessageGraph;
use crate::config::{StoreType, DEFAULT_ID_HEADER};
use crate::{Error, Result};
use crate::{Settings, DEFAULT_CALLBACK_URL_HEADER_KEY};

/// [Message] with a oneshot for notifying when the message has been completed processed.
pub(crate) struct MessageWrapper {
    // TODO: this might be more that saving to ISB.
    pub(crate) confirm_save: oneshot::Sender<()>,
    pub(crate) message: Message,
}

/// Serving payload passed on to Numaflow.
#[derive(Debug)]
pub struct Message {
    pub value: Bytes,
    pub id: String,
    pub headers: HashMap<String, String>,
}

enum ActorMessage {
    Read {
        batch_size: usize,
        timeout_at: Instant,
        reply_to: oneshot::Sender<Result<Vec<Message>>>,
    },
    Ack {
        offsets: Vec<String>,
        reply_to: oneshot::Sender<Result<()>>,
    },
}

/// Background actor that starts Axum server for accepting HTTP requests.
struct ServingSourceActor {
    /// The HTTP handlers will put the message received from the payload to this channel
    messages: mpsc::Receiver<MessageWrapper>,
    /// Channel for the actor handle to communicate with this actor
    handler_rx: mpsc::Receiver<ActorMessage>,
    /// Mapping from request's ID header (usually `X-Numaflow-Id` header) to a channel.
    /// Sending a message on this channel notifies the HTTP handler function that the message
    /// has been successfully processed.
    tracker: HashMap<String, oneshot::Sender<()>>,
    vertex_replica_id: u16,
    callback_url: String,
}

impl ServingSourceActor {
    async fn start(
        js_context: Context,
        settings: Arc<Settings>,
        handler_rx: mpsc::Receiver<ActorMessage>,
        request_channel_buffer_size: usize,
        vertex_replica_id: u16,
    ) -> Result<()> {
        // Channel to which HTTP handlers will send request payload
        let (messages_tx, messages_rx) = mpsc::channel(request_channel_buffer_size);
        // create a callback store for tracking
        let callback_store =
            JetStreamCallbackStore::new(js_context.clone(), &settings.js_store).await?;
        // Create the message graph from the pipeline spec and the redis store
        let msg_graph = MessageGraph::from_pipeline(&settings.pipeline_spec).map_err(|e| {
            Error::InitError(format!(
                "Creating message graph from pipeline spec: {:?}",
                e
            ))
        })?;

        let callback_url = format!(
            "https://{}:{}/v1/process/callback",
            &settings.host_ip, &settings.app_listen_port
        );

        // Create a redis store to store the callbacks and the custom responses
        match &settings.store_type {
            StoreType::Nats => {
                let nats_store = JetStreamDataStore::new(js_context, &settings.js_store).await?;
                let callback_state =
                    CallbackState::new(msg_graph, nats_store, callback_store).await?;
                let app = crate::AppState {
                    message: messages_tx,
                    settings,
                    orchestrator_state: callback_state,
                };
                tokio::spawn(async move {
                    crate::serve(app).await.unwrap();
                });
            }
            StoreType::UserDefined(ud_config) => {
                let ud_store = UserDefinedStore::new(ud_config.clone()).await?;
                let callback_state =
                    CallbackState::new(msg_graph, ud_store, callback_store).await?;
                let app = crate::AppState {
                    message: messages_tx,
                    settings,
                    orchestrator_state: callback_state,
                };
                tokio::spawn(async move {
                    crate::serve(app).await.unwrap();
                });
            }
        }

        tokio::spawn(async move {
            let mut serving_actor = ServingSourceActor {
                messages: messages_rx,
                handler_rx,
                tracker: HashMap::new(),
                vertex_replica_id,
                callback_url,
            };
            serving_actor.run().await;
        });

        Ok(())
    }

    async fn run(&mut self) {
        while let Some(msg) = self.handler_rx.recv().await {
            self.handle_message(msg).await;
        }
    }

    async fn handle_message(&mut self, actor_msg: ActorMessage) {
        match actor_msg {
            ActorMessage::Read {
                batch_size,
                timeout_at,
                reply_to,
            } => {
                let messages = self.read(batch_size, timeout_at).await;
                let _ = reply_to.send(messages);
            }
            ActorMessage::Ack { offsets, reply_to } => {
                let status = self.ack(offsets).await;
                let _ = reply_to.send(status);
            }
        }
    }

    // for monovertex after reading register the oneshot in the callback handler
    async fn read(&mut self, count: usize, timeout_at: Instant) -> Result<Vec<Message>> {
        let mut messages = vec![];
        loop {
            // Stop if the read timeout has reached or if we have collected the requested number of messages
            if messages.len() >= count || Instant::now() >= timeout_at {
                break;
            }
            let next_msg = self.messages.recv();
            let message = match tokio::time::timeout_at(timeout_at, next_msg).await {
                Ok(Some(msg)) => msg,
                Ok(None) => {
                    // If we have collected at-least one message, we return those messages.
                    // The error will happen on all the subsequent read attempts too.
                    if messages.is_empty() {
                        return Err(Error::Other(
                            "Sending half of the Serving channel has disconnected".into(),
                        ));
                    }
                    tracing::error!("Sending half of the Serving channel has disconnected");
                    return Ok(messages);
                }
                Err(_) => return Ok(messages),
            };
            let MessageWrapper {
                confirm_save,
                mut message,
            } = message;

            self.tracker.insert(message.id.clone(), confirm_save);
            message.headers.insert(
                DEFAULT_CALLBACK_URL_HEADER_KEY.into(),
                self.callback_url.clone(),
            );
            message
                .headers
                .insert(DEFAULT_ID_HEADER.into(), message.id.clone());
            messages.push(message);
        }
        Ok(messages)
    }

    async fn ack(&mut self, offsets: Vec<String>) -> Result<()> {
        let offset_suffix = format!("-{}", self.vertex_replica_id);
        for offset in offsets {
            let offset = offset.strip_suffix(&offset_suffix).ok_or_else(|| {
                Error::Source(format!("offset does not end with '{}'", &offset_suffix))
            })?;
            let confirm_save_tx = self
                .tracker
                .remove(offset)
                .ok_or_else(|| Error::Source("offset was not found in the tracker".into()))?;
            confirm_save_tx
                .send(())
                .map_err(|e| Error::Source(format!("Sending on confirm_save channel: {e:?}")))?;
        }
        Ok(())
    }
}

#[derive(Clone)]
pub struct ServingSource {
    batch_size: usize,
    // timeout for each batch read request
    timeout: Duration,
    actor_tx: mpsc::Sender<ActorMessage>,
}

impl ServingSource {
    pub async fn new(
        context: Context,
        settings: Arc<Settings>,
        batch_size: usize,
        timeout: Duration,
        vertex_replica_id: u16,
    ) -> Result<Self> {
        let (actor_tx, actor_rx) = mpsc::channel(2 * batch_size);
        ServingSourceActor::start(
            context,
            settings,
            actor_rx,
            2 * batch_size,
            vertex_replica_id,
        )
        .await?;
        Ok(Self {
            batch_size,
            timeout,
            actor_tx,
        })
    }

    pub async fn read_messages(&self) -> Result<Vec<Message>> {
        let start = Instant::now();
        let (tx, rx) = oneshot::channel();
        let actor_msg = ActorMessage::Read {
            reply_to: tx,
            batch_size: self.batch_size,
            timeout_at: Instant::now() + self.timeout,
        };
        let _ = self.actor_tx.send(actor_msg).await;
        let messages = rx.await.map_err(Error::ActorTaskTerminated)??;
        tracing::debug!(
            count = messages.len(),
            requested_count = self.batch_size,
            time_taken_ms = start.elapsed().as_millis(),
            "Got messages from Serving source"
        );
        Ok(messages)
    }

    pub async fn ack_messages(&self, offsets: Vec<String>) -> Result<()> {
        let (tx, rx) = oneshot::channel();
        let actor_msg = ActorMessage::Ack {
            offsets,
            reply_to: tx,
        };
        let _ = self.actor_tx.send(actor_msg).await;
        rx.await.map_err(Error::ActorTaskTerminated)??;
        Ok(())
    }
}

#[cfg(test)]
mod tests {
    use std::{sync::Arc, time::Duration};

    use async_nats::jetstream;

    use super::ServingSource;
    use crate::Settings;

    type Result<T> = std::result::Result<T, Box<dyn std::error::Error>>;

    #[cfg(feature = "nats-tests")]
    #[tokio::test]
    async fn test_serving_source() -> Result<()> {
        let js_url = "localhost:4222";
        let client = async_nats::connect(js_url).await.unwrap();
        let context = jetstream::new(client);
        let store_name = "test_serving_source";

        let _ = context.delete_key_value(store_name).await;
        context
            .create_key_value(jetstream::kv::Config {
                bucket: store_name.to_string(),
                history: 5,
                ..Default::default()
            })
            .await
            .unwrap();

        // Set up the CryptoProvider (controls core cryptography used by rustls) for the process
        let _ = rustls::crypto::aws_lc_rs::default_provider().install_default();

        let settings = Arc::new(Settings {
            js_store: store_name.to_string(),
            app_listen_port: 2444,
            ..Default::default()
        });

        let serving_source = ServingSource::new(
            context,
            Arc::clone(&settings),
            10,
            Duration::from_millis(1),
            0,
        )
        .await?;

        let client = reqwest::Client::builder()
            .timeout(Duration::from_secs(2))
            .danger_accept_invalid_certs(true)
            .build()
            .unwrap();

        // Wait for the server
        for _ in 0..10 {
            let resp = client
                .get(format!(
                    "https://localhost:{}/livez",
                    settings.app_listen_port
                ))
                .send()
                .await;
            if resp.is_ok() {
                break;
            }
            tokio::time::sleep(Duration::from_millis(10)).await;
        }

        tokio::spawn(async move {
            loop {
                tokio::time::sleep(Duration::from_millis(10)).await;
                let mut messages = serving_source.read_messages().await.unwrap();
                if messages.is_empty() {
                    // Server has not received any requests yet
                    continue;
                }
                assert_eq!(messages.len(), 1);
                let msg = messages.remove(0);
                serving_source
                    .ack_messages(vec![format!("{}-0", msg.id)])
                    .await
                    .unwrap();
                break;
            }
        });

        let resp = client
            .post(format!(
                "https://localhost:{}/v1/process/async",
                settings.app_listen_port
            ))
            .json("test-payload")
            .send()
            .await?;

        assert!(resp.status().is_success());
        Ok(())
    }
}
