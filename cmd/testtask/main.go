package main

import (
	"log"

	"github.com/exceedololo/notion-/internal/config"
	"github.com/exceedololo/notion-/internal/transport/http"
)

func main() {
	cfg := config.LoadConfig()

	server := http.NewServer(cfg)
	err := server.Run()
	if err != nil {
		log.Fatal(err)
	}
}
