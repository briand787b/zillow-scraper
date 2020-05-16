package controller

import (
	"zcrapr/core/model"
	"zcrapr/core/plog"
)

// PropertyController controls the flow of HTTP routes for Property resources
type PropertyController struct {
	l  plog.Logger
	ps model.PropertyStore
}

// NewPropertyController returns a new PropertyController
func NewPropertyController(l plog.Logger, ps model.PropertyStore) *PropertyController {
	return &PropertyController{
		l:  l,
		ps: ps,
	}
}
