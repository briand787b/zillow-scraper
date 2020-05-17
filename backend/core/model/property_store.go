package model

import "context"

// PropertyStore is anything that can store Propertyes
type PropertyStore interface {
	GetAllCapturesByPropertyID(ctx context.Context, propID string) ([]Capture, error)
	GetAllPropertyIDs(ctx context.Context, skip, take int) ([]string, error)
	GetLatestCaptureByPropertyID(ctx context.Context, propID string) (*Capture, error)
	GetPropertyByID(ctx context.Context, id string) (*Property, error)
	InsertCaptureByPropertyID(ctx context.Context, propID string, c *Capture) error
	InsertProperty(ctx context.Context, p *Property) error
}
