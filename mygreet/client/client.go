package client

import (
	"context"
	"log"
	"strconv"

	"github.com/golang/protobuf/ptypes/empty"

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
		conn, err := grpc.Dial(conf.GrpcAddr, grpc.WithInsecure())
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
	conn, err := grpc.Dial(conf.GrpcAddr, grpc.WithTransportCredentials(creds))
	if err != nil {
		panic(err)
	}
	c.client = v1.NewGreetServiceClient(conn)
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

	for i := 0; i < 3; i++ {
		req := &v1.GreetRequest{
			Greeting: &v1.Greeting{
				FirstName: "Ivo " + strconv.Itoa(i+1),
				LastName:  "Stoyanov",
			},
		}
		resp, err := c.client.Greet(ctx, req)
		if err != nil {
			log.Printf("*** error: %+v", err)
			continue
		}
		log.Printf("Greet response: %+v", resp)
	}
	return nil
}
