package server

import (
	"log"
	"net"

	"google.golang.org/grpc/credentials"

	"github.com/ivostoyanov-bc/sandbox/mygreet/config"
	"github.com/ivostoyanov-bc/sandbox/mygreet/protos"
	"google.golang.org/grpc"
)

type Server struct {
	conf *config.Config
	srv  *grpc.Server
}

func New(conf *config.Config) *Server {
	if conf.Secure == 0 {
		s := &Server{
			conf: conf,
			srv:  grpc.NewServer(),
		}
		protos.RegisterGreetServiceServer(s.srv, s)
		return s
	}
	creds, err := credentials.NewServerTLSFromFile(conf.CertFile, conf.KeyFile)
	if err != nil {
		panic(err)
	}
	s := &Server{
		conf: conf,
		srv:  grpc.NewServer(grpc.Creds(creds)),
	}
	protos.RegisterGreetServiceServer(s.srv, s)
	return s
}

func (s *Server) ListenAndServe() error {
	l, err := net.Listen("tcp", s.conf.Endpoint)
	if err != nil {
		return err
	}
	log.Printf("grpc Server ListenAndServe on %v, secure: %v", s.conf.Endpoint, s.conf.Secure)
	err = s.srv.Serve(l)
	return err
}
