package http

import (
	"github.com/exceedololo/notion-/internal/transport/http/handlers"
	"net/http"
)

func RegisterHandlers(router *http.ServeMux) {
	router.HandleFunc("/redis/incr", handlers.IncrementHandler)
}
