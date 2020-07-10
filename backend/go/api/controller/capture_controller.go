package controller

import (
	"net/http"
	"zcrapr/core/model"
	"zcrapr/core/perr"
	"zcrapr/core/plog"

	"github.com/go-chi/render"
)

// CaptureController controls the flow of HTTP routes for Capture resources
type CaptureController struct {
	l  plog.Logger
	cs model.CaptureStore
	ps model.PropertyStore
}

// NewCaptureController returns a new CaptureController
func NewCaptureController(l plog.Logger, cs model.CaptureStore, ps model.PropertyStore) *CaptureController {
	return &CaptureController{
		l:  l,
		cs: cs,
		ps: ps,
	}
}

// HandleCreate handles the creation of Captures
func (c *CaptureController) HandleCreate(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var data CaptureRequest
	if err := render.Bind(r, &data); err != nil {
		render.Render(w, r, perr.NewValidationHTTPErrorFromError(ctx, err, "could not bind request", c.l))
		return
	}

	if err := data.c.Save(ctx, c.l, data.p, c.cs, c.ps); err != nil {
		c.l.Error(ctx, "failed to create a capture record",
			"request", data,
			"error", err,
		)
		render.Render(w, r, perr.NewHTTPErrorFromError(ctx, err, "could not save Capture", c.l))
		return
	}

	render.Status(r, http.StatusCreated)
}
