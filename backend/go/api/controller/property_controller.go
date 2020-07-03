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

func (c *PropertyController) propertyCtx(next http.Handler) http.Handler {
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

// HandleCreate handles the creation of a Property
func (c *PropertyController) HandleCreate(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var data PropertyRequest
	if err := render.Bind(r, &data); err != nil {
		render.Render(w, r, perr.NewValidationHTTPErrorFromError(ctx, err, "could not bind request", c.l))
		return
	}

	if err := data.p.Save(ctx, c.l, c.ps); err != nil {
		render.Render(w, r, perr.NewHTTPErrorFromError(ctx, err, "could not persist Property", c.l))
		return
	}

	render.Status(r, http.StatusCreated)
	render.Render(w, r, NewPropertyResponse(data.p))
}

// HandleGetByID handles the retrieval of a Property by its ID
func (c *PropertyController) HandleGetByID(w http.ResponseWriter, r *http.Request) {
	m, ok := r.Context().Value(propertyCtxKey).(*model.Property)
	if !ok {
		render.Render(w, r, perr.NewNotFoundHTTPError(r.Context(), "no or invalid Property value for propertyCtxKey", c.l))
		return
	}

	render.Render(w, r, NewPropertyResponse(m))
}

// HandleGetAll handles requests to get all the properties
func (c *PropertyController) HandleGetAll(w http.ResponseWriter, r *http.Request) {
	if addr := r.URL.Query().Get("address"); addr != "" {
		ctx := r.Context()
		props, err := model.GetPropertiesByAddress(ctx, c.l, addr, c.ps)
		if err != nil {
			render.Render(w, r, perr.NewHTTPErrorFromError(ctx, err, "could not get property by address", c.l))
			return
		}

		render.Render(w, r, NewPropertyResponseList(props))
		return
	}

	// skip := r.Context().Value(skipCtxKey).(int)
	take := r.Context().Value(takeCtxKey).(int)

	ctx := r.Context()
	props, err := model.GetAllProperties(ctx, c.l, take, c.ps)
	if err != nil {
		render.Render(w, r, perr.NewHTTPErrorFromError(ctx, err, "could not get all root Media", c.l))
		return
	}

	render.Render(w, r, NewPropertyResponseList(props))
}
