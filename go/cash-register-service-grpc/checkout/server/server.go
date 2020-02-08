package server

import (
	"context"
	"log"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	v1 "github.com/lt-schmidt-jr/cash-register-service-IVOSTOYANOV/shared/api/checkout"
	"github.com/lt-schmidt-jr/cash-register-service-IVOSTOYANOV/shared/pkg/config"
	"github.com/lt-schmidt-jr/cash-register-service-IVOSTOYANOV/shared/pkg/version"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	// the simplest catalog ever - key is the UPC -> value is the item
	catalog map[string]v1.Item
	mutex   sync.Mutex
	conf    *config.Config
	srv     *grpc.Server
	mux     *http.ServeMux
}

func New(conf *config.Config) (s *Server) {
	switch conf.Secure {
	case 0:
		s = &Server{
			conf: conf,
			srv:  grpc.NewServer(),
		}
	case 1:
		// with TLS
		creds, err := credentials.NewServerTLSFromFile(conf.CertFile, conf.KeyFile)
		if err != nil {
			panic(err)
		}
		s = &Server{
			conf: conf,
			srv:  grpc.NewServer(grpc.Creds(creds)),
		}
	case 2:
		// todo
	}
	version.Name = "checkout"
	// Register reflection service to enable grpc introspection
	reflection.Register(s.srv)
	v1.RegisterCheckoutServiceServer(s.srv, s)
	s.catalog = make(map[string]v1.Item)
	return s
}

func (s *Server) ListenAndServe() error {
	var err error

	l, err := net.Listen("tcp", s.conf.GrpcAddr)
	if err != nil {
		return err
	}

	log.Printf("%s gRPC Server on %v, secure: %v", version.Name, s.conf.GrpcAddr, s.conf.Secure)

	go func() {
		err = s.srv.Serve(l)
		if err != nil {
			log.Printf("grpc serve error %v", err)
		}
	}()
	// no ssl for REST yet
	// todo: secure
	opts := []grpc.DialOption{grpc.WithInsecure()}
	mux := runtime.NewServeMux()

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	time.Sleep(1 * time.Second)
	// REST endpoint
	err = v1.RegisterCheckoutServiceHandlerFromEndpoint(ctx, mux, s.conf.GrpcAddr, opts)
	if err != nil {
		log.Printf("Register service error %v", err)
		return err
	}
	log.Printf("%s REST Server on %v, secure: %v", version.Name, s.conf.RestAddr, s.conf.Secure)
	return http.ListenAndServe(s.conf.RestAddr, mux)
}
