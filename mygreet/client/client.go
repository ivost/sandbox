package client

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/ivostoyanov-bc/sandbox/mygreet/config"
	"github.com/ivostoyanov-bc/sandbox/mygreet/protos"
)

type Client struct {
	conf   *config.Config
	client protos.GreetServiceClient
}

func New(conf *config.Config) *Client {
	c := &Client{conf: conf}
	if conf.Secure == 0 {
		conn, err := grpc.Dial(conf.Endpoint, grpc.WithInsecure())
		if err != nil {
			panic(err)
		}
		c.client = protos.NewGreetServiceClient(conn)
		return c
	}
	creds, err := credentials.NewClientTLSFromFile(conf.CertFile, "")
	if err != nil {
		panic(err)
	}
	conn, err := grpc.Dial(conf.Endpoint, grpc.WithTransportCredentials(creds))
	if err != nil {
		panic(err)
	}
	c.client = protos.NewGreetServiceClient(conn)
	return c
}

func (c *Client) DoUnary() error {
	req := &protos.GreetRequest{
		Greeting: &protos.Greeting{
			FirstName: "Ivo",
			LastName:  "Stoyanov",
		},
	}
	ctx := context.Background()
	// unary
	resp, err := c.client.Greet(ctx, req)
	log.Printf("resp %+v, err %v", resp, err)
	return err
}
