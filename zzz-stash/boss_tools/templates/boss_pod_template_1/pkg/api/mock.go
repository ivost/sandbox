package api

import (
	"time"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func NewMockServer() *Server {
	config := &Config{
		Port:                      "8080",
		HttpServerShutdownTimeout: 5 * time.Second,
		HttpServerTimeout:         30 * time.Second,
		BackendURL:                "",
		ConfigPath:                "/config",
		DataPath:                  "/data",
		HttpClientTimeout:         30 * time.Second,
		UIColor:                   "blue",
		UIPath:                    ".ui",
		UIMessage:                 "Greetings",
		Hostname:                  "localhost",
	}

	logger, _ := zap.NewDevelopment()

	return &Server{
		router: mux.NewRouter(),
		logger: logger,
		config: config,
	}
}
