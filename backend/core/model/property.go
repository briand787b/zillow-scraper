package model

import (
	"context"

	"github.com/pkg/errors"
)

// Status enumerates the statuses a property can be in
type Status int

const (
	// ForSale is any property that is for sale
	ForSale Status = iota
	// Pending is a property under contract
	Pending
	// Sold is a property that is off the market
	Sold
)

// Capture is a set of values captured at a specific point in time
type Capture struct {
	Price   uint   `json:"price"`
	Acreage uint   `json:"acrege"`
	Status  Status `json:"status"`
}

// Property represents a Property
type Property struct {
	ID  string
	URL string

	captures []Capture
}

// Add adds a new Capture to a Property
func (p *Property) Add(ctx context.Context, ps PropertyStore, c *Capture) error {
	return errors.New("NOT IMPLEMENTED")
}

// // GetCapturesShallow retrieves all the already loaded Captures
// func (p *Property) GetCapturesShallow() []Capture {
// 	return p.captures
// }

// func (p *Property)
