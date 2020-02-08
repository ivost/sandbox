package server

import (
	"fmt"
	"log"
	"time"

	v1 "github.com/lt-schmidt-jr/cash-register-service-IVOSTOYANOV/shared/api/checkout"
)

// here is the main "business" logic to build the receipt by grouping items
// input: array with UPCs from the scan - assume in scan order
// output: result is pointer to CheckoutResponse which includes Receipt
// fails fast by returning error if UPC is not in the catalog
func (s *Server) makeReceipt(register string, items []*v1.ScanItem, result *v1.CheckoutResponse) error {
	log.Printf("makeReceipt input: %+v", items)
	result.Receipt = &v1.Receipt{
		Register: register,
		Time:     time.Now().Unix(),
	}
	// step 1 - group by UPC
	// use map as intermediate structure for line items
	lines := make(map[string]v1.LineItem)
	for _, code := range items {
		line, found := lines[code.Upc]
		if found {
			// UPC found - increment (no need to lookup)
			line.Qty++
			line.Extprice += line.Price
			lines[code.Upc] = line
			continue
		}
		// not in map - lookup item, add
		it := s.lookup(code.Upc)
		if it == nil {
			return fmt.Errorf("UPC %v not found", code.Upc)
		}
		// first line
		lines[code.Upc] = v1.LineItem{
			Id:       code.Upc,
			Qty:      1,
			Name:     it.Name,
			Price:    it.Price,
			Extprice: it.Price,
		}
	}
	// step 2 - build the receipt - must be in chronological order
	// iterate original list to keep the order
	result.Receipt.Items = make([]*v1.LineItem, len(lines))
	idx := 0
	for _, scan := range items {
		if line, found := lines[scan.Upc]; found {
			result.Receipt.Items[idx] = &line
			result.Receipt.Total += line.Extprice
			idx++
			// remove from map to shortcut other scans with same UPC
			delete(lines, scan.Upc)
		}
	}
	return nil
}

// find item by UPC
func (s *Server) lookup(UPC string) *v1.Item {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	it, ok := s.catalog[UPC]
	if !ok {
		return nil
	}
	return &it
}
