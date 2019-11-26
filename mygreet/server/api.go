package server

import (
	"context"
	"io"
	"log"
	"math/rand"

	"github.com/ivostoyanov-bc/sandbox/mygreet/protos"
)

func (s *Server) Greet(ctx context.Context, req *protos.GreetRequest) (resp *protos.GreetResponse, err error) {
	log.Printf("Greet %+v", req)
	resp = &protos.GreetResponse{
		Result: "Hello " + req.GetGreeting().GetFirstName(),
	}
	return
}

func (s *Server) ServerStream(req *protos.ServerStreamRequest, stream protos.GreetService_ServerStreamServer) error {
	p := int64(1)
	max := req.GetMaxPrime()
	for ; p < max; p += 2 {
		if isPrime(p) {
			resp := &protos.ServerStreamResponse{
				Prime: p,
			}
			stream.Send(resp)
		}
	}
	return nil
}

func (s *Server) ClientStream(stream protos.GreetService_ClientStreamServer) error {
	resp := &protos.ClientStreamResponse{
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

func (s *Server) BiDirStream(stream protos.GreetService_BiDirStreamServer) error {

	resp := &protos.BiDirStreamResponse{StartIndex: -1, EndIndex: -1}

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
