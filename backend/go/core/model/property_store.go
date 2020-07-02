package model

import "context"

// PropertyStore is anything that can store Propertyes
type PropertyStore interface {
	GetAllPropertyIDs(ctx context.Context, take int) ([]string, error)
	GetPropertiesByAddress(ctx context.Context, address string) ([]Property, error)
	GetPropertyByID(ctx context.Context, id string) (*Property, error)
	GetPropertyIDByAddress(ctx context.Context, address string) (string, error)
	GetPropertyIDByURL(ctx context.Context, url string) (string, error)
	InsertProperty(ctx context.Context, p *Property) error
	UpdateProperty(ctx context.Context, p *Property) error

	// TODO
	// convert GetPropertyByID to GetPropertiesByIDs to work better w/ GetAllPropertyIDs
	// FUTURE
	// experiment with putting Property IDs behind a sorted set - may have better
	// performace when searching (index-like), though it does require a 2nd degree
	// of searching (sorted-set -> propID -> prop hash)
}
