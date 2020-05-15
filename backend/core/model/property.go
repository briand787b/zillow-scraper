package model

import "context"

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
	Price   uint
	Acreage uint
	Status
}

// Property represents a Property
type Property struct {
	Captures []Capture
}

// Add adds a new Capture to a Property
func (p *Property) Add(ctx context.Context, as PropertyStore, c *Capture) error {

}
