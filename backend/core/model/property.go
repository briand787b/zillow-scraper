package model

import (
	"context"
	"fmt"
	"zcrapr/core/perr"

	"github.com/pkg/errors"
)

const maxPropertyTake = 100

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
	Price   int    `json:"price"`
	Acreage int    `json:"acrege"`
	Status  Status `json:"status"`
}

// Property represents a Property
type Property struct {
	ID  string
	URL string

	captures []Capture
}

// GetPropertyByID gets a Property by its ID
func GetPropertyByID(ctx context.Context, id string, ps PropertyStore) (*Property, error) {
	if id == "" {
		return nil, perr.NewErrInvalid("cannot search with an empty ID")
	}

	p, err := ps.GetPropertyByID(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, "could not get property by ID from store")
	}

	return p, nil
}

// GetAllProperties retrieves all Properties
func GetAllProperties(ctx context.Context, skip, take int, ps PropertyStore) ([]Property, error) {
	if skip < 0 {
		return nil, perr.NewErrInvalid("skip cannot be negative")
	}

	if take < 1 {
		return nil, perr.NewErrInvalid("take must be at least 1")
	}

	if take > maxPropertyTake {
		return nil, perr.NewErrInvalid(fmt.Sprintf("cannot take more than max (%v)", maxPropertyTake))
	}

	propIDs, err := ps.GetAllPropertyIDs(ctx, skip, take)
	if err != nil {
		return nil, errors.Wrap(err, "could not get all properties")
	}

	props := make([]Property, len(propIDs))
	for i, propID := range propIDs {
		prop, err := ps.GetPropertyByID(ctx, propID)
		if err != nil {
			return nil, errors.Wrap(err, "could not get property by ID")
		}

		props[i] = *prop
	}

	return props, nil
}

// Save saves a property to the database
func (p *Property) Save(ctx context.Context, ps PropertyStore) error {
	if p.URL == "" {
		return perr.NewErrInvalid("URL cannot be empty")
	}

	if p.ID == "" {
		if err := ps.InsertProperty(ctx, p); err != nil {
			return errors.Wrap(err, "could not insert Property")
		}

		return nil
	}

	if err := ps.UpdateProperty(ctx, p); err != nil {
		return errors.Wrap(err, "could not update Property")
	}

	return nil
}

// AddCapture adds a new Capture to a Property
func (p *Property) AddCapture(ctx context.Context, c *Capture, ps PropertyStore) error {
	if c.Acreage < 1 {
		return perr.NewErrInvalid("properties must have at least one acre")
	}

	if c.Price < 1 {
		return perr.NewErrInvalid("nothing in life is free")
	}

	switch c.Status {
	case ForSale, Pending, Sold:
	default:
		return perr.NewErrInvalid("capture has invalid state")
	}

	if err := ps.InsertCaptureByPropertyID(ctx, p.ID, c); err != nil {
		return errors.Wrap(err, "could not insert property")
	}

	return nil
}

// GetCaptures retrieves all the already loaded Captures
func (p *Property) GetCaptures() []Capture {
	return p.captures
}

// LoadCaptures loads all captures into the Property receiever
func (p *Property) LoadCaptures(ctx context.Context, ps PropertyStore) error {
	caps, err := ps.GetAllCapturesByPropertyID(ctx, p.ID)
	if err != nil {
		return errors.Wrap(err, "could not get captures")
	}

	p.captures = caps
	return nil
}
