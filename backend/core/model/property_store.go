package model

import "context"

// PropertyStore is anything that can store Propertyes
type PropertyStore interface {
	GetAllCapturesByPropertyID(ctx context.Context, propID string) ([]Capture, error)
	GetAllPropertyIDs(ctx context.Context, skip, take int) ([]string, error)
	GetLatestCaptureByPropertyID(ctx context.Context, propID string) (*Capture, error)
	GetPropertyByID(ctx context.Context, id string) (*Property, error)
	GetPropertyByURL(ctx context.Context, url string) (*Property, error)
	InsertCaptureByPropertyID(ctx context.Context, propID string, c *Capture) error
	InsertProperty(ctx context.Context, p *Property) error
	UpdateProperty(ctx context.Context, p *Property) error

	// TODO
	// GetPropertyByURL(ctx context.Context, url string)
	// convert GetPropertyByID to GetPropertiesByIDs to work better w/ GetAllPropertyIDs
	// FUTURE
	// experiment with putting Property IDs behind a sorted set - may have better
	// performace when searching (index-like), though it does require a 2nd degree
	// of searching (sorted-set -> propID -> prop hash)
}
