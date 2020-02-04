package client

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	v1 "github.com/lt-schmidt-jr/cash-register-service-IVOSTOYANOV/shared/api/checkout"
	"github.com/lt-schmidt-jr/cash-register-service-IVOSTOYANOV/shared/pkg/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type Client struct {
	conf   *config.Config
	client v1.CheckoutServiceClient
}

func New(conf *config.Config) *Client {
	if conf == nil {
		conf = config.DefaultConfig()
	}
	c := &Client{conf: conf}
	if conf.Secure == 0 {
		conn, err := grpc.Dial(conf.GrpcAddr, grpc.WithInsecure())
		if err != nil {
			panic(err)
		}
		c.client = v1.NewCheckoutServiceClient(conn)
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
	c.client = v1.NewCheckoutServiceClient(conn)
	return c
}

func (c *Client) Health() (*v1.HealthResponse, error) {
	ctx := context.Background()
	return c.client.Health(ctx, &empty.Empty{})
}
