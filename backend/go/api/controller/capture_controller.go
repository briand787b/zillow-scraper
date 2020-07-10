package controller

import (
	"context"
	"net/http"
	"zcrapr/core/model"
	"zcrapr/core/perr"
	"zcrapr/core/plog"

	"github.com/go-chi/chi"
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

func (c *CaptureController) propertyCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		pID := chi.URLParam(r, "property_id")
		if pID == "" {
			render.Render(w, r, perr.NewInternalServerHTTPError(ctx, "no property_id in url params", c.l))
			return
		}

		p, err := model.GetPropertyByID(r.Context(), c.l, pID, c.ps)
		if err != nil {
			render.Render(w, r, perr.NewHTTPErrorFromError(ctx, err, "could not find Property by ID", c.l))
			return
		}

		ctx = context.WithValue(r.Context(), propertyCtxKey, p)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
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

// HandleGetAllByPropertyID gets all captures by property ID
func (c *CaptureController) HandleGetAllByPropertyID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	property, ok := ctx.Value(propertyCtxKey).(*model.Property)
	if !ok {
		render.Render(w, r, perr.NewNotFoundHTTPError(ctx, "could not find property by ID", c.l))
	}

	caps, err := model.GetAllCapturesByPropertyID(ctx, c.l, property.ID, c.cs)
	if err != nil {
		render.Render(w, r, perr.NewHTTPErrorFromError(ctx, err, "could not get all captures by propertyID", c.l))
		return
	}

	render.Render(w, r, NewCaptureResponseList(caps))
}
