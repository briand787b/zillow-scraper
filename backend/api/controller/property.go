package controller

import (
	"net/http"

	"zcrapr/core/model"
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

	ms []model.Property
}

// NewPropertyResponseList converts a slice of model.Property into a PropertyResponseList
func NewPropertyResponseList(mms []model.Property, skip, take int) *PropertyResponseList {
	return &PropertyResponseList{
		Skip: skip,
		Take: take,

		ms: mms,
	}
}

// Render does any processing ahead of the go-chi library's rendering
func (l *PropertyResponseList) Render(w http.ResponseWriter, r *http.Request) error {
	l.Properties = make([]PropertyResponse, len(l.ms))
	for i := 0; i < len(l.ms); i++ {
		l.Properties[i] = *NewPropertyResponse(&l.ms[i])
	}

	if len(l.Properties) >= l.Take {
		l.NextSkip = l.Skip + l.Take
	}

	return nil
}
