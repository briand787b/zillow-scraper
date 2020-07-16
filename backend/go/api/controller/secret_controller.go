package controller

import (
	"net/http"
	"zcrapr/core/plog"

	"github.com/go-chi/render"
)

type SecretController struct {
	l                     plog.Logger
	googleMapsEmbedAPIKey string
}

func NewSecretController(l plog.Logger, googleMapsEmbedAPIKey string) *SecretController {
	return &SecretController{
		l:                     l,
		googleMapsEmbedAPIKey: googleMapsEmbedAPIKey,
	}
}

func (c *SecretController) HandleGetGoogleMapsEmbedSecret(w http.ResponseWriter, r *http.Request) {
	render.Render(w, r, NewSecretResponse(c.googleMapsEmbedAPIKey))
}
