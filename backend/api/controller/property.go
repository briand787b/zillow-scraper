package controller

import (
	"net/http"

	"zcrape/core/model"
)

// PropertyRequest is a request object for the Property resource
type PropertyRequest struct {
}

// Bind does processing on the PropertyRequest after it gets decoded
func (m *PropertyRequest) Bind(r *http.Request) error {
	// m.Encoding = obj.Encoding(strings.ToLower(string(m.Encoding)))
	return nil
}

// Property converts a PropertyRequest to a model.Property
func (m *PropertyRequest) Property() *model.Property {
	// mm := model.Property{
	// 	ID:       m.ID,
	// 	Name:     m.Name,
	// 	Length:   m.Length,
	// 	Encoding: m.Encoding,
	// }

	// for _, c := range m.Children {
	// 	mm.Children = append(mm.Children, *c.Property())
	// }

	// return &mm

	return nil
}

// PropertyResponse represents the response object for Property requests
type PropertyResponse struct {
	model.Property
}

// NewPropertyResponse creates a new PropertyResponse
func NewPropertyResponse(mm *model.Property) *PropertyResponse {
	return &PropertyResponse{*mm}
}

// Render processes a PropertyResponse before rendering in HTTP response
func (m *PropertyResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// PropertyResponseList represents a list of Property
type PropertyResponseList struct {
	Property []model.Property `json:"Property"`
	Skip     int              `json:"skip"`
	Take     int              `json:"take"`
	NextSkip int              `json:"next_skip,omitempty"`
}

// NewPropertyResponseList converts a slice of model.Property into a PropertyResponseList
func NewPropertyResponseList(mms []model.Property, skip, take int) *PropertyResponseList {
	return &PropertyResponseList{
		Property: mms,
		Skip:     skip,
		Take:     take,
	}
}

// Render does any processing ahead of the go-chi library's rendering
func (l *PropertyResponseList) Render(w http.ResponseWriter, r *http.Request) error {
	if len(l.Property) >= l.Take {
		l.NextSkip = l.Skip + l.Take
	}

	return nil
}
