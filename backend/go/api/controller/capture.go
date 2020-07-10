package controller

import (
	"net/http"

	"zcrapr/core/model"
	"zcrapr/core/perr"

	"github.com/pkg/errors"
)

// CaptureRequest is a request object for the Capture resource
type CaptureRequest struct {
	Property PropertyRequest `json:"property"`
	Price    int             `json:"price"`
	Status   string          `json:"status"`

	c *model.Capture
	p *model.Property
}

// Bind does processing on the CaptureRequest after it gets decoded
func (c *CaptureRequest) Bind(r *http.Request) error {
	status, err := model.NewStatus(c.Status)
	if err != nil {
		return errors.Wrap(err, "could not get valid status from request")
	}

	c.c = &model.Capture{
		Price:  c.Price,
		Status: status,
	}

	if err := c.Property.Bind(r); err != nil {
		return errors.Wrap(err, "could not bind Property")
	}

	c.p = c.Property.p

	return nil
}

// CaptureResponse represents the response object for Capture requests
type CaptureResponse struct {
	Price  int    `json:"price"`
	Status string `json:"Status"`

	c *model.Capture
}

// NewCaptureResponse creates a new CaptureResponse
func NewCaptureResponse(mc *model.Capture) *CaptureResponse {
	return &CaptureResponse{
		c: mc,
	}
}

// Render processes a CaptureResponse before rendering in HTTP response
func (c *CaptureResponse) Render(w http.ResponseWriter, r *http.Request) error {
	c.Price = c.c.Price
	c.Status = c.c.Status.String()

	return nil
}

// CaptureResponseList represents a list of Capture
type CaptureResponseList struct {
	Captures []CaptureResponse `json:"captures"`

	cs []model.Capture
}

// NewCaptureResponseList converts a slice of model.Capture into a CaptureResponseList
func NewCaptureResponseList(mcs []model.Capture) *CaptureResponseList {
	return &CaptureResponseList{cs: mcs}
}

// Render does any processing ahead of the go-chi library's rendering
func (l *CaptureResponseList) Render(w http.ResponseWriter, r *http.Request) error {
	l.Captures = make([]CaptureResponse, len(l.cs))
	for i := 0; i < len(l.cs); i++ {
		l.Captures[i] = *NewCaptureResponse(&l.cs[i])
		if err := l.Captures[i].Render(nil, nil); err != nil {
			return perr.NewErrInternal(errors.Wrap(err, "could not bind PropertyResponse"))
		}
	}

	return nil
}
