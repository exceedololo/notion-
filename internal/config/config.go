package config

import (
	"log"
	"os"
)

type Config struct {
	Host string
	Port string
}

func LoadConfig() *Config {
	host := os.Getenv("HOST")
	if host == "" {
		host = "localhost"
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Loaded configuration: Host=%s, Port=%s", host, port)

	return &Config{
		Host: host,
		Port: port,
	}
}
