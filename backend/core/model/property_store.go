package model

import "context"

// PropertyStore is anything that can store Propertyes
type PropertyStore interface {
	AddCapture(ctx context.Context, url string, c *Capture) error
	GetLatestCapture(ctx context.Context, url string) (*Capture, error)
}
