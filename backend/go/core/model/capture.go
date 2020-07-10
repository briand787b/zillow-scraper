package model

import (
	"context"
	"zcrapr/core/perr"
	"zcrapr/core/plog"

	"github.com/pkg/errors"
)

// Status enumerates the statuses a property can be in
type Status int

const (
	// ForSale is any property that is for sale
	ForSale Status = iota + 1
	// Pending is a property under contract
	Pending
	// Sold is a property that is off the market
	Sold
	// OffMarket is a property that is not on the market
	OffMarket
)

// NewStatus returns a new Status from a status string
func NewStatus(statusString string) (Status, error) {
	switch statusString {
	case "For Sale", "for sale", "ForSale":
		return ForSale, nil
	case "Pending", "pending", "Under Contract":
		return Pending, nil
	case "Sold", "sold":
		return Sold, nil
	case "Off Market", "OffMarket", "off market":
		return OffMarket, nil
	default:
		return Status(0), perr.NewErrInvalid("provided string is not a valid Status string")
	}
}

func (s *Status) String() string {
	switch *s {
	case ForSale:
		return "For Sale"
	case Pending:
		return "Pending"
	case Sold:
		return "Sold"
	case OffMarket:
		return "Off Market"
	default:
		return "Invalid"
	}
}

// Capture is a set of values captured at a specific point in time
type Capture struct {
	Price  int    `json:"price"`
	Status Status `json:"status"`
}

// GetAllCapturesByPropertyID retrieves all Captures by Property ID
func GetAllCapturesByPropertyID(ctx context.Context, l plog.Logger, propertyID string, cs CaptureStore) ([]Capture, error) {
	if propertyID == "" {
		return nil, perr.NewErrInvalid("propertyID is empty string")
	}

	caps, err := cs.GetAllCapturesByPropertyID(ctx, propertyID)
	if err != nil {
		return nil, errors.Wrap(err, "could not get captures from CaptureStore")
	}

	return caps, nil
}

// Save saves a Capture to the persitence layer.  It is expected to
func (c *Capture) Save(ctx context.Context, l plog.Logger, p *Property, cs CaptureStore, ps PropertyStore) error {
	if c.Price < 1 {
		return perr.NewErrInvalid("nothing in life is free")
	}

	switch c.Status {
	case ForSale, Pending, Sold, OffMarket:
	default:
		return perr.NewErrInvalid("capture has invalid state")
	}

	propFromDB, err := LoadPropertyByFields(ctx, l, p, ps)
	if err != nil {
		// the only acceptable error here is a NotFound. Even then it must not have
		// an ID, as this implies its searching for a known Property that clearly does
		// not exist
		if !perr.SameType(err, perr.ErrNotFound) || p.ID != "" {
			return errors.Wrap(err, "LoadPropertyByFields failed with unacceptable error")
		}

		// Property does not exist, try to create it
		l.Info(ctx, "Property provided to Capture.Save does not exit, attempting to create it",
			"Capture", c,
			"Property", p,
		)

		if err := p.Save(ctx, l, ps); err != nil {
			return errors.Wrap(err, "could not save new Property to database")
		}

		propFromDB = p
	}

	if err := cs.InsertCaptureByPropertyID(ctx, propFromDB.ID, c); err != nil {
		return errors.Wrap(err, "could not insert Capture")
	}

	return nil
}
