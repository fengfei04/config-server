package router

import (
	"config-server/internal/handler"

	"github.com/gin-gonic/gin"
)

type Router struct {
	engine *gin.Engine
}

func New() *Router {
	return &Router{
		engine: gin.Default(),
	}
}

func (r *Router) InitRoute() {
	groupV1 := r.engine.Group("/v1")
	{
		groupV1.GET("/namespaces/:namespace/:kind", handler.GetResources)
		groupV1.POST("/namespaces/:namespace/:kind", handler.CreateResource)
		groupV1.GET("/namespaces/:namespace/:kind/:name", handler.GetResource)
		groupV1.PATCH("/namespaces/:namespace/:kind/:name", handler.UpdateResource)
		groupV1.DELETE("/namespaces/:namespace/:kind/:name", handler.DeleteResource)
	}
}

func (r *Router) Run(addr ...string) {
	r.engine.Run(addr...)
}
