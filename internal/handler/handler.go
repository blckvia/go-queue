package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"

	"gtihub.com/blckvia/go-queue/internal/broker"

	"github.com/gin-contrib/pprof"
)

type Handler struct {
	broker *broker.Broker
}

func New(broker *broker.Broker) *Handler {
	return &Handler{
		broker: broker,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.Use(otelgin.Middleware("go-queue"))

	debugMode := viper.GetBool("debug")
	if debugMode {
		pprof.Register(router, "/debug/pprof")
	}

	v1 := router.Group("/v1/queues/:queue_name")
	{
		v1.POST("/messages", h.sendMessage)
		v1.POST("/subscriptions", h.subscribe)
	}

	return router
}
