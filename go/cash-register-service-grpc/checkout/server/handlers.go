package server

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	v1 "github.com/lt-schmidt-jr/cash-register-service-IVOSTOYANOV/shared/api/checkout"
	"github.com/lt-schmidt-jr/cash-register-service-IVOSTOYANOV/shared/pkg/system"
	"github.com/lt-schmidt-jr/cash-register-service-IVOSTOYANOV/shared/pkg/version"
)

func (s *Server) Health(ctx context.Context, none *empty.Empty) (resp *v1.HealthResponse, err error) {
	resp = &v1.HealthResponse{
		Name:    "checkout",
		Version: version.Version,
		Build:   version.Build,
		Status:  "OK",
		Time:    time.Now().Format(time.RFC3339),
		Address: system.MyIP(),
	}
	return
}

// add/replace catalog item
func (s *Server) AddItem(ctx context.Context, req *v1.AddItemRequest) (*v1.AddItemResponse, error) {
	it := req.Item
	if it == nil {
		return nil, fmt.Errorf("no item")
	}
	log.Printf("item %+v", it)
	// simple data validation
	id := it.Id
	if id == "" || it.Name == "" || it.Price <= 0 {
		return nil, fmt.Errorf("invalid item")
	}
	it.Valid = true
	s.mutex.Lock()
	defer s.mutex.Unlock()
	// upsert
	s.catalog[id] = *it
	return &v1.AddItemResponse{
		Item: it,
	}, nil
}

// perform checkout
// input is array of scanned items
// outtput is the receipt
func (s *Server) Checkout(ctx context.Context, req *v1.CheckoutRequest) (*v1.CheckoutResponse, error) {
	var resp v1.CheckoutResponse
	err := s.makeReceipt(req.Register, req.Items, &resp)
	return &resp, err
}
