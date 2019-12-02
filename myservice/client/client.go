package client

import (
	"context"
	"log"

	"github.com/golang/protobuf/ptypes/empty"

	"github.com/ivost/sandbox/myservice/config"
	v1 "github.com/ivost/sandbox/myservice/myservice"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type Client struct {
	conf   *config.Config
	client v1.MyServiceClient
}

func New(conf *config.Config) *Client {
	c := &Client{conf: conf}
	if conf.Secure == 0 {
		conn, err := grpc.Dial(conf.GrpcAddr, grpc.WithInsecure())
		if err != nil {
			panic(err)
		}
		c.client = v1.NewMyServiceClient(conn)
		return c
	}
	creds, err := credentials.NewClientTLSFromFile(conf.CertFile, "")
	if err != nil {
		panic(err)
	}
	conn, err := grpc.Dial(conf.GrpcAddr, grpc.WithTransportCredentials(creds))
	if err != nil {
		panic(err)
	}
	c.client = v1.NewMyServiceClient(conn)
	return c
}

func (c *Client) DoUnary() error {
	ctx := context.Background()
	// unary
	log.Printf("Health check")
	hr, err := c.client.Health(ctx, &empty.Empty{})
	if err != nil {
		log.Printf("*** error: %+v", err)
	} else {
		log.Printf("Health response: %+v", hr)
	}
	return err
}
