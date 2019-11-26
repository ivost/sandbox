package api

import (
	"context"
	"net/http"
	"time"

	"github.com/braincorp/boss_pod_template/pkg/version"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

var ()

type Config struct {
	VersionFlag               bool          `mapstructure:"version"`
	CloneFlag                 bool          `mapstructure:"clone"`
	BackendURL                string        `mapstructure:"backend-url"`
	DataPath                  string        `mapstructure:"data-path"`
	ConfigPath                string        `mapstructure:"config-path"`
	HostPort                  string        `mapstructure:"host-port"`
	HttpClientTimeout         time.Duration `mapstructure:"http-client-timeout"`
	HttpServerTimeout         time.Duration `mapstructure:"http-server-timeout"`
	HttpServerShutdownTimeout time.Duration `mapstructure:"http-server-shutdown-timeout"`
}

type Server struct {
	router *mux.Router
	logger *zap.Logger
	config *Config
}

func NewServer(config *Config, logger *zap.Logger) (*Server, error) {
	srv := &Server{
		router: mux.NewRouter(),
		logger: logger,
		config: config,
	}

	return srv, nil
}

func (s *Server) registerMiddlewares() {
	httpLogger := NewLoggingMiddleware(s.logger)
	s.router.Use(httpLogger.Handler)
	s.router.Use(versionMiddleware)
}

func (s *Server) ListenAndServe(stopCh <-chan struct{}) {

	s.setupRoutes()
	s.registerMiddlewares()

	srv := &http.Server{
		Addr:         s.config.HostPort,
		WriteTimeout: s.config.HttpServerTimeout,
		ReadTimeout:  s.config.HttpServerTimeout,
		IdleTimeout:  10 * s.config.HttpServerTimeout,
		Handler:      s.router,
	}

	// run the server in the background
	go func() {
		s.logger.Sugar().Infof("Server %v %v listens on %+v",
			version.VERSION, version.REVISION, s.config.HostPort)
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			s.logger.Fatal("HTTP server crashed", zap.Error(err))
		}
	}()

	// wait for SIGTERM or SIGINT
	<-stopCh
	ctx, cancel := context.WithTimeout(context.Background(), s.config.HttpServerShutdownTimeout)
	defer cancel()

	// attempt graceful shutdown
	if err := srv.Shutdown(ctx); err != nil {
		s.logger.Warn("HTTP server graceful shutdown failed", zap.Error(err))
	} else {
		s.logger.Info("HTTP server stopped")
	}
}
