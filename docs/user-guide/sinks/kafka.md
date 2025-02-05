# Kafka Sink

Two methods are available for integrating Kafka topics into your Numaflow pipeline:
using a user-defined Kafka Sink or opting for the built-in Kafka Sink provided by Numaflow.

## Option 1: User-Defined Kafka Sink

Developed and maintained by the Numaflow contributor community,
the [Kafka Sink](https://github.com/numaproj-contrib/kafka-java) provides a reliable and feature-complete solution for publishing messages to Kafka topics.

Key Features:

* **Customization:** Offers complete control over Kafka Sink configurations to tailor to specific requirements.
* **Kafka Java Client Utilization:** Leverages the Kafka Java client for reliable message publishing to Kafka topics.
* **Schema Management:** Integrates seamlessly with the Confluent Schema Registry to support schema validation and manage schema evolution effectively.

More details on how to use the Kafka Sink can be found [here](https://github.com/numaproj-contrib/kafka-java?tab=readme-ov-file#write-data-to-kafka).

## Option 2: Built-in Kafka Sink

A `Kafka` sink is used to forward the messages to a Kafka topic. Kafka sink supports configuration overrides.

### Kafka Headers

We will insert `keys` into the Kafka header, but since `keys` is an array, we will add `keys` into the header in the
following format.

* `__keys_len` will have the number of `key` in the header. if `__keys_len` == `0`, means no `keys` are present.
* `__keys_%d` will have the `key`, e.g., `__key_0` will be the first key, and so forth.

### Example 

```yaml
spec:
  vertices:
    - name: kafka-output
      sink:
        kafka:
          brokers:
            - my-broker1:19700
            - my-broker2:19700
          topic: my-topic
          tls: # Optional.
            insecureSkipVerify: # Optional, where to skip TLS verification. Default to false.
            caCertSecret: # Optional, a secret reference, which contains the CA Cert.
              name: my-ca-cert
              key: my-ca-cert-key
            certSecret: # Optional, pointing to a secret reference which contains the Cert.
              name: my-cert
              key: my-cert-key
            keySecret: # Optional, pointing to a secret reference which contains the Private Key.
              name: my-pk
              key: my-pk-key
          sasl: # Optional
            mechanism: GSSAPI # PLAIN, GSSAPI, OAUTHBEARER, SCRAM-SHA-256 or SCRAM-SHA-512 other mechanisms not supported
            gssapi: # Optional, for GSSAPI mechanism
              serviceName: my-service
              realm: my-realm
              # KRB5_USER_AUTH for auth using password
              # KRB5_KEYTAB_AUTH for auth using keytab
              authType: KRB5_KEYTAB_AUTH
              usernameSecret: # Pointing to a secret reference which contains the username
                name: gssapi-username
                key: gssapi-username-key
              # Pointing to a secret reference which contains the keytab (authType: KRB5_KEYTAB_AUTH)
              keytabSecret:
                name: gssapi-keytab
                key: gssapi-keytab-key
              # Pointing to a secret reference which contains the keytab (authType: KRB5_USER_AUTH)
              passwordSecret:
                name: gssapi-password
                key: gssapi-password-key
              kerberosConfigSecret: # Pointing to a secret reference which contains the kerberos config
                name: my-kerberos-config
                key: my-kerberos-config-key
            plain: # Optional, for PLAIN mechanism
              userSecret: # Pointing to a secret reference which contains the user
                name: plain-user
                key: plain-user-key
              passwordSecret: # Pointing to a secret reference which contains the password
                name: plain-password
                key: plain-password-key
              # Send the Kafka SASL handshake first if enabled (defaults to true)
              # Set this to false if using a non-Kafka SASL proxy
              handshake: true
            scramsha256: # Optional, for SCRAM-SHA-256 mechanism
              userSecret: # Pointing to a secret reference which contains the user
                name: scram-sha-256-user
                key: scram-sha-256-user-key
              passwordSecret: # Pointing to a secret reference which contains the password
                name: scram-sha-256-password
                key: scram-sha-256-password-key
              # Send the Kafka SASL handshake first if enabled (defaults to true)
              # Set this to false if using a non-Kafka SASL proxy
              handshake: true 
            scramsha512: # Optional, for SCRAM-SHA-512 mechanism
              userSecret: # Pointing to a secret reference which contains the user
                name: scram-sha-512-user
                key: scram-sha-512-user-key
              passwordSecret: # Pointing to a secret reference which contains the password
                name: scram-sha-512-password
                key: scram-sha-512-password-key
              # Send the Kafka SASL handshake first if enabled (defaults to true)
              # Set this to false if using a non-Kafka SASL proxy
              handshake: true 
            oauth:  #Optional, for OAUTHBEARER mechanism
              clientID: # Pointing to a secret reference which contains the client id
                name: kafka-oauth-client
                key: clientid 
              clientSecret: # Pointing to a secret reference which contains the client secret
                name: kafka-oauth-client
                key: clientsecret 
              tokenEndpoint: https://oauth-token.com/v1/token
          # Optional, a yaml format string which could apply more configuration for the sink.
          # The configuration hierarchy follows the Struct of sarama.Config at https://github.com/IBM/sarama/blob/main/config.go.
          config: |
            producer:
            compression: 2
```
