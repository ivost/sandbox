package client

import (
	"context"
	"io"
	"log"
	"time"

	"github.com/ivostoyanov-bc/sandbox/mygreet/protos"
)

func (c *Client) DoServerStream() error {
	ctx := context.Background()
	sreq := &protos.ServerStreamRequest{
		MaxPrime: 1000,
	}
	stream, err := c.client.ServerStream(ctx, sreq)
	if err != nil {
		return err
	}

	for {
		resp, err := stream.Recv()
		switch err {
		case nil:
			log.Printf("resp: %+v", resp)
		case io.EOF:
			log.Printf("EOF")
			stream.CloseSend()
			return nil
		default:
			log.Printf("Error %v", err)
			return err
		}
	}
}

func (c *Client) doClientStream() error {
	ctx := context.Background()
	stream, err := c.client.ClientStream(ctx)
	if err != nil {
		return err
	}

	for i := int64(1); i <= 10; i++ {
		req := &protos.ClientStreamRequest{
			Index:  i,
			Number: i * 10,
		}
		err = stream.Send(req)
		if err != nil {
			return err
		}
	}
	resp, err := stream.CloseAndRecv()
	if err != nil {
		return err
	}
	log.Printf("response %+v", resp)
	return err
}

func doBiDirStream(client protos.GreetServiceClient) error {
	ctx := context.Background()
	stream, err := client.BiDirStream(ctx)
	if err != nil {
		return err
	}

	done := make(chan struct{})
	// send in goroutine
	go func() {
		max := int64(1000000)
		for i := int64(1); i <= max; i++ {
			req := &protos.BiDirStreamRequest{
				Index:  i,
				Number: i * 2,
			}
			err = stream.Send(req)
			if err != nil {
				return
			}
			time.Sleep(100 * time.Millisecond)
		}
		stream.CloseSend()
	}()
	// recv in goroutine
	go func() {
		for {
			res, err := stream.Recv()
			if err != nil {
				close(done)
				return
			}
			log.Printf("received %+v", res)
		}
	}()

	// block
	<-done
	return nil
}
