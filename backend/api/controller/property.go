package controller

import (
	"net/http"

	"zcrapr/core/model"
	"zcrapr/core/perr"

	"github.com/pkg/errors"
)

// PropertyRequest is a request object for the Property resource
type PropertyRequest struct {
	ID  string `json:"id"`
	URL string `json:"url"`

	p *model.Property
}

// Bind does processing on the PropertyRequest after it gets decoded
func (m *PropertyRequest) Bind(r *http.Request) error {
	m.p = &model.Property{
		ID:  m.ID,
		URL: m.URL,
	}

	return nil
}

// PropertyResponse represents the response object for Property requests
type PropertyResponse struct {
	ID  string `json:"id"`
	URL string `json:"url"`

	p *model.Property
}

// NewPropertyResponse creates a new PropertyResponse
func NewPropertyResponse(mp *model.Property) *PropertyResponse {
	return &PropertyResponse{
		p: mp,
	}
}

// Render processes a PropertyResponse before rendering in HTTP response
func (m *PropertyResponse) Render(w http.ResponseWriter, r *http.Request) error {
	m.ID = m.p.ID
	m.URL = m.p.URL

	return nil
}

// PropertyResponseList represents a list of Property
type PropertyResponseList struct {
	Properties []PropertyResponse `json:"properties"`
	Skip       int                `json:"skip"`
	Take       int                `json:"take"`
	NextSkip   int                `json:"next_skip,omitempty"`

	ps []model.Property
}

// NewPropertyResponseList converts a slice of model.Property into a PropertyResponseList
func NewPropertyResponseList(mps []model.Property, skip, take int) *PropertyResponseList {
	return &PropertyResponseList{
		Skip: skip,
		Take: take,

		ps: mps,
	}
}

// Render does any processing ahead of the go-chi library's rendering
func (l *PropertyResponseList) Render(w http.ResponseWriter, r *http.Request) error {
	l.Properties = make([]PropertyResponse, len(l.ps))
	for i := 0; i < len(l.ps); i++ {
		l.Properties[i] = *NewPropertyResponse(&l.ps[i])
		if err := l.Properties[i].Render(nil, nil); err != nil {
			return perr.NewErrInternal(errors.Wrap(err, "could not bind PropertyResponse"))
		}
	}

	if len(l.Properties) >= l.Take {
		l.NextSkip = l.Skip + l.Take
	}

	return nil
}
