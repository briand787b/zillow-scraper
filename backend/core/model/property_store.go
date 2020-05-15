package model

// PropertyStore is anything that can store Propertyes
type PropertyStore interface {
	AddCapture(a *Property, c *Capture) error
	GetCaptures(a *Property) ([]Capture, error)
}
