package http

import (
	"github.com/exceedololo/notion-/internal/transport/http/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterHandlers() *gin.Engine {
	router := gin.Default()

	router.POST("/redis/incr", handlers.IncrementHandler)

	return router
}
