package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/numaproj/numaflow/server/apis/v1_1"
)

func v1_1Routes(r gin.IRouter) {
	handler, err := v1_1.NewHandler()
	if err != nil {
		panic(err)
	}
	// List all namespaces that have Pipeline or InterStepBufferService objects.
	r.GET("/namespaces", handler.ListNamespaces)
	// Summarized information of all the namespaces in a cluster wrapped in a list.
	r.GET("/cluster-summary")
	// Create a Pipeline.
	r.POST("/namespaces/:namespace/pipelines")
	// All pipelines for a given namespace.
	r.GET("/namespaces/:namespace/pipelines")
	// Get a Pipeline information.
	r.GET("/namespaces/:namespace/pipelines/:pipeline")
	// Update a Pipeline.
	r.PUT("/namespaces/:namespace/pipelines/:pipeline")
	// Delete a Pipeline.
	r.DELETE("/namespaces/:namespace/pipelines/:pipeline")
	// Patch the pipeline spec to achieve operations such as "pause" and "resume".
	r.PATCH("/namespaces/:namespace/pipelines/:pipeline")
	// Create an InterStepBufferService object.
	r.POST("/namespaces/:namespace/isb-services")
	// List all the InterStepBufferService objects for a given namespace.
	r.GET("/namespaces/:namespace/isb-services")
	// Get an InterStepBufferService object.
	r.GET("/namespaces/:namespace/isb-services/:isb-services")
	// Update an InterStepBufferService object.
	r.PUT("/namespaces/:namespace/isb-services/:isb-services")
	// Delete an InterStepBufferService object.
	r.DELETE("/namespaces/:namespace/isb-services/:isb-services")
	// Get all the Inter-Step Buffers of a pipeline.
	r.GET("/namespaces/:namespace/pipelines/:pipeline/isbs")
	// Get all the watermarks information of a pipeline.
	r.GET("/namespaces/:namespace/pipelines/:pipeline/watermarks")
	// Get a vertex information of a pipeline. TODO: do we need it?
	r.GET("/namespaces/:namespace/pipelines/:pipeline/vertices/:vertex")
	// Update a vertex spec.
	r.PUT("/namespaces/:namespace/pipelines/:pipeline/vertices/:vertex")
	// Get all the vertex metrics of a pipeline. TODO: to be finalized
	r.GET("/namespaces/:namespace/pipelines/:pipeline/vertices/:vertex/metrics")
	// Get all the pods of a vertex.
	r.GET("/namespaces/:namespace/pipelines/:pipeline/vertices/:vertex/pods")
	// Get the metrics such as cpu, memory usage for a pod.
	r.GET("/metrics/namespaces/:namespace/pods/:pod")
	// Get pod logs.
	r.GET("/namespaces/:namespace/pods/:pod/logs")
	// List of the Kubernetes events of a namespace.
	r.GET("/namespaces/:namespace/events")

}