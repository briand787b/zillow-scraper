package model

import "context"

// CaptureStore is anything that can store/retrieve Captures
type CaptureStore interface {
	GetAllCapturesByPropertyID(ctx context.Context, propID string) ([]Capture, error)
	GetLatestCaptureByPropertyID(ctx context.Context, propID string) (*Capture, error)
	InsertCaptureByPropertyID(ctx context.Context, propID string, c *Capture) error
}
