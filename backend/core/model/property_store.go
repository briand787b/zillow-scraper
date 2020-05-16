package model

import "context"

// PropertyStore is anything that can store Propertyes
type PropertyStore interface {
	AddCapture(ctx context.Context, a *Property, c *Capture) error
	GetCaptures(ctx context.Context, a *Property) ([]Capture, error)
}
