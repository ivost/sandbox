package client

import (
	"context"
	"log"

	"github.com/ivost/sandbox/mygreet/config"
	v1 "github.com/ivost/sandbox/mygreet/greet"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type Client struct {
	conf   *config.Config
	client v1.GreetServiceClient
}

func New(conf *config.Config) *Client {
	c := &Client{conf: conf}
	if conf.Secure == 0 {
		conn, err := grpc.Dial(conf.Endpoint, grpc.WithInsecure())
		if err != nil {
			panic(err)
		}
		c.client = v1.NewGreetServiceClient(conn)
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
	c.client = v1.NewGreetServiceClient(conn)
	return c
}

func (c *Client) DoUnary() error {
	req := &v1.GreetRequest{
		Greeting: &v1.Greeting{
			FirstName: "Ivo",
			LastName:  "Stoyanov",
		},
	}
	ctx := context.Background()
	// unary
	for i := 0; i < 5; i++ {
		resp, err := c.client.Greet(ctx, req)
		if err != nil {
			log.Printf("*** error: %+v", err)
			continue
		}
		log.Printf("response: %+v", resp)
	}
	return nil
}
