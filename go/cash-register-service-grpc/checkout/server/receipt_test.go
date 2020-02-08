package server

import (
	"context"
	"log"
	"net"
	"testing"
	"time"

	v1 "github.com/lt-schmidt-jr/cash-register-service-IVOSTOYANOV/shared/api/checkout"
	"github.com/lt-schmidt-jr/cash-register-service-IVOSTOYANOV/shared/pkg/config"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

const (
	// some scan codes
	UPC1 = "111"
	UPC2 = "222"
	// some register
	REG1 = "1"
	// for buffcon - which allows us testing of the full gRPC stack
	bufSize = 1024 * 1024
)

var lis *bufconn.Listener

func init() {
	lis = bufconn.Listen(bufSize)
	s := New(config.DefaultConfig())
	go func() {
		if err := s.srv.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func getClient(t *testing.T, ctx context.Context) v1.CheckoutServiceClient {
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	//defer conn.Close()
	return v1.NewCheckoutServiceClient(conn)

}

func bufDialer(string, time.Duration) (net.Conn, error) {
	return lis.Dial()
}

func TestAddItems(t *testing.T) {
	addItem(t, v1.Item{Id: UPC1, Name: "foo", Price: 20})
	addItem(t, v1.Item{Id: UPC2, Name: "bar", Price: 20})
	// upsert
	addItem(t, v1.Item{Id: UPC1, Name: "foo", Price: 10})
}

func TestReceipt(t *testing.T) {
	TestAddItems(t)
	req := &v1.CheckoutRequest{Register: "1", Items: make([]*v1.ScanItem, 0)}
	req.Items = append(req.Items, &v1.ScanItem{Upc: UPC1})
	req.Items = append(req.Items, &v1.ScanItem{Upc: UPC2})
	req.Items = append(req.Items, &v1.ScanItem{Upc: UPC1})
	ctx := context.Background()
	client := getClient(t, ctx)
	resp, err := client.Checkout(ctx, req)
	assert.NoError(t, err)
	r := resp.Receipt
	assert.EqualValues(t, 2, len(r.Items))
	assert.EqualValues(t, 40, r.Total)
	// first receipt line
	l := r.Items[0]
	assert.EqualValues(t, UPC1, l.Id)
	assert.EqualValues(t, 2, l.Qty)
	assert.EqualValues(t, 20, l.Extprice)
	// next
	l = r.Items[1]
	assert.EqualValues(t, UPC2, l.Id)
	assert.EqualValues(t, 1, l.Qty)
	assert.EqualValues(t, 20, l.Extprice)
}

func TestReceiptBadScan(t *testing.T) {
	TestAddItems(t)
	req := &v1.CheckoutRequest{Register: "1", Items: make([]*v1.ScanItem, 0)}
	req.Items = append(req.Items, &v1.ScanItem{Upc: UPC1})
	req.Items = append(req.Items, &v1.ScanItem{Upc: "999999"})
	ctx := context.Background()
	client := getClient(t, ctx)
	_, err := client.Checkout(ctx, req)
	assert.Error(t, err)
}

func addItem(t *testing.T, item v1.Item) {
	ctx := context.Background()
	client := getClient(t, ctx)
	req := &v1.AddItemRequest{Item: &item}
	resp, err := client.AddItem(ctx, req)
	assert.NoError(t, err)
	assert.True(t, resp.Item.Valid)
	assert.Equal(t, req.Item.Price, resp.Item.Price)
	assert.Equal(t, req.Item.Name, resp.Item.Name)
	assert.Equal(t, req.Item.Id, resp.Item.Id)
}
