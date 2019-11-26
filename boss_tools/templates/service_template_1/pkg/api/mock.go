package api

import (
	"time"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func NewMockServer() *Server {
	config := &Config{
		HostPort:                  ":8080",
		HttpServerShutdownTimeout: 5 * time.Second,
		HttpServerTimeout:         15 * time.Second,
		HttpClientTimeout:         15 * time.Second,
		BackendURL:                "",
		ConfigPath:                "/config",
		DataPath:                  "/data",
	}

	logger, _ := zap.NewDevelopment()

	return &Server{
		router: mux.NewRouter(),
		logger: logger,
		config: config,
	}
}
