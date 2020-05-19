package controller

import (
	"net/http"

	"zcrapr/core/model"
	"zcrapr/core/perr"

	"github.com/pkg/errors"
)

// CaptureRequest is a request object for the Capture resource
type CaptureRequest struct {
	Price   int    `json:"price"`
	Status  string `json:"status"`
	URL     string `json:"url"`
	Address string `json:"address"`

	c *model.Capture
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
	Properties []CaptureResponse `json:"properties"`
	Skip       int               `json:"skip"`
	Take       int               `json:"take"`
	NextSkip   int               `json:"next_skip,omitempty"`

	ps []model.Capture
}

// NewCaptureResponseList converts a slice of model.Capture into a CaptureResponseList
func NewCaptureResponseList(mps []model.Capture, skip, take int) *CaptureResponseList {
	return &CaptureResponseList{
		Skip: skip,
		Take: take,

		ps: mps,
	}
}

// Render does any processing ahead of the go-chi library's rendering
func (l *CaptureResponseList) Render(w http.ResponseWriter, r *http.Request) error {
	l.Properties = make([]CaptureResponse, len(l.ps))
	for i := 0; i < len(l.ps); i++ {
		l.Properties[i] = *NewCaptureResponse(&l.ps[i])
		if err := l.Properties[i].Render(nil, nil); err != nil {
			return perr.NewErrInternal(errors.Wrap(err, "could not bind CaptureResponse"))
		}
	}

	if len(l.Properties) >= l.Take {
		l.NextSkip = l.Skip + l.Take
	}

	return nil
}
