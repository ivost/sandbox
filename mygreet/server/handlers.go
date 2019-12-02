package server

import (
	"context"
	"io"
	"log"
	"math/rand"
	"time"

	"github.com/golang/protobuf/ptypes/empty"

	v1 "github.com/ivost/sandbox/mygreet/greet"
)

func (s *Server) Health(ctx context.Context, none *empty.Empty) (resp *v1.HealthResponse, err error) {
	resp = &v1.HealthResponse{
		Status: "OK",
		Time:   time.Now().Format(time.RFC3339),
	}
	return
}

func (s *Server) Greet(ctx context.Context, req *v1.GreetRequest) (resp *v1.GreetResponse, err error) {
	log.Printf("Received Greet request: %+v", req)
	resp = &v1.GreetResponse{
		Result: "Hello, " + req.GetGreeting().GetFirstName(),
	}
	return
}

func (s *Server) ServerStream(req *v1.ServerStreamRequest, stream v1.GreetService_ServerStreamServer) error {
	p := int64(1)
	max := req.GetMaxPrime()
	for ; p < max; p += 2 {
		if isPrime(p) {
			resp := &v1.ServerStreamResponse{
				Prime: p,
			}
			stream.Send(resp)
		}
	}
	return nil
}

func (s *Server) ClientStream(stream v1.GreetService_ClientStreamServer) error {
	resp := &v1.ClientStreamResponse{
		StartIndex: -1,
		EndIndex:   -1,
		Sum:        0,
	}

	for {
		req, err := stream.Recv()
		switch err {
		case nil:
			log.Printf("req: %+v", req)
			if resp.StartIndex < 0 {
				resp.StartIndex = req.Index
			}
			resp.EndIndex = req.Index
			resp.Sum += req.Number
		case io.EOF:
			log.Printf("EOF")
			stream.SendAndClose(resp)
			return nil
		default:
			log.Printf("Error %v", err)
			stream.SendAndClose(resp)
			return err
		}
	}
}

func (s *Server) BiDirStream(stream v1.GreetService_BiDirStreamServer) error {

	resp := &v1.BiDirStreamResponse{StartIndex: -1, EndIndex: -1}

	// pick random mod 1-10
	randNum := 1 + rand.Int31n(9)
	idx := int32(0)
	for {
		// send response at random index
		if idx%randNum == 0 {
			if idx > 0 {
				log.Printf("sending: %+v", resp)
				err := stream.Send(resp)
				if err != nil {
					return err
				}
			}
			resp.StartIndex = -1
			resp.EndIndex = -1
			resp.Sum = 0
		}
		idx++
		// receive from client
		req, err := stream.Recv()
		switch err {
		case nil:
			log.Printf("received from client: %+v", req)
			if resp.StartIndex < 0 {
				resp.StartIndex = req.Index
			}
			resp.EndIndex = req.Index
			resp.Sum += req.Number
		case io.EOF:
			log.Printf("EOF")
			stream.Send(resp)
			return nil
		default:
			log.Printf("Error %v", err)
			stream.Send(resp)
			return err
		}
	}
}
