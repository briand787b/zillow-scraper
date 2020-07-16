package controller

import "net/http"

// SecretResponse is the response type for Secret requests
type SecretResponse struct {
	Secret string `json:"secret"`
}

func NewSecretResponse(secret string) *SecretResponse {
	return &SecretResponse{
		Secret: secret,
	}
}

func (r *SecretResponse) Render(w http.ResponseWriter, rq *http.Request) error {
	return nil
}
