package server

import (
	"context"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"

	v1 "github.com/ivost/sandbox/mygreet/greet"

	"github.com/ivost/sandbox/mygreet/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	conf *config.Config
	srv  *grpc.Server
	mux  *http.ServeMux
}

func New(conf *config.Config) *Server {

	// Register gRPC server endpoint
	if conf.Secure == 0 {
		s := &Server{
			conf: conf,
			srv:  grpc.NewServer(),
		}

		// Register reflection service on gRPC server.
		reflection.Register(s.srv)
		v1.RegisterGreetServiceServer(s.srv, s)
		return s
	}

	// with TLS
	creds, err := credentials.NewServerTLSFromFile(conf.CertFile, conf.KeyFile)
	if err != nil {
		panic(err)
	}

	s := &Server{
		conf: conf,
		srv:  grpc.NewServer(grpc.Creds(creds)),
	}

	// Register reflection service on gRPC server.
	reflection.Register(s.srv)

	v1.RegisterGreetServiceServer(s.srv, s)

	return s
}

func (s *Server) ListenAndServe() error {
	var err error

	l, err := net.Listen("tcp", s.conf.GrpcAddr)
	if err != nil {
		return err
	}

	log.Printf("grpc Server ListenAndServe on %v, secure: %v", s.conf.GrpcAddr, s.conf.Secure)

	go func() {
		err = s.srv.Serve(l)
		if err != nil {
			log.Printf("grpc serve error %v", err)
		}
	}()

	// https://grpc-ecosystem.github.io/grpc-gateway/docs/usage.html
	// Note: Make sure the gRPC server is running properly and accessible
	opts := []grpc.DialOption{grpc.WithInsecure()}
	mux := runtime.NewServeMux()

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	time.Sleep(1 * time.Second)
	err = v1.RegisterGreetServiceHandlerFromEndpoint(ctx, mux, s.conf.GrpcAddr, opts)
	if err != nil {
		log.Printf("Register service error %v", err)
		return err
	}
	log.Printf("REST Server ListenAndServe on %v/greet, secure: %v", 8080, s.conf.Secure)
	return http.ListenAndServe(":8080", mux)
	return err
}
