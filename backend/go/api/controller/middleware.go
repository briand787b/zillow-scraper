package controller

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"zcrapr/core/perr"
	"zcrapr/core/plog"

	"github.com/go-chi/render"
	"github.com/pkg/errors"
)

const (
	// CorsEnvVarKey is the key used to get the CORS host
	CorsEnvVarKey = "PL_CORS_HOST"
)

// Middleware acts as a bridge between the request and the controllers
type Middleware struct {
	corsHost string
	l        plog.Logger
	uuidGen  fmt.Stringer
}

// NewMiddleware returns a new Middleware
func NewMiddleware(l plog.Logger, uuidGen fmt.Stringer, corsHost string) (*Middleware, error) {
	if l == nil {
		return nil, errors.New("cannot have a nil logger")
	}

	if uuidGen == nil {
		return nil, errors.New("cannot have a nil uuidGen")
	}

	if corsHost == "" {
		l.Error(context.Background(), "corsHost is empty string")
	}

	return &Middleware{l: l, uuidGen: uuidGen, corsHost: corsHost}, nil
}

func (m *Middleware) spanAndTrace(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := plog.StoreSpanIDTraceID(r.Context(), m.uuidGen.String(), m.uuidGen.String())
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (m *Middleware) logRoute(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		m.l.Info(r.Context(), "started handling HTTP request",
			"uri", r.RequestURI,
			"method", r.Method,
		)
		next.ServeHTTP(w, r)
	})
}

func (m *Middleware) disableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", m.corsHost)
		w.Header().Set("Access-Control-Allow-Methods", "POST,GET,PUT,DELETE,OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "*")

		next.ServeHTTP(w, r)
	})
}

func (m *Middleware) skipTake(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s, err := getSkip(r)
		if err != nil {
			render.Render(w, r, perr.NewValidationHTTPErrorFromError(r.Context(), err, "could not get skip", m.l)) // NOTE: THIS IS NOT THE CORRECT RESPONSE, JUST TESTING...
			return
		}

		t, err := getTake(r)
		if err != nil {
			render.Render(w, r, perr.NewValidationHTTPErrorFromError(r.Context(), err, "could not get take", m.l)) // NOTE: THIS IS NOT THE CORRECT RESPONSE, JUST TESTING...
			return
		}

		ctx := context.WithValue(r.Context(), skipCtxKey, s)
		ctx = context.WithValue(ctx, takeCtxKey, t)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// getSkip returns the value of the skip query string parameter
// It defaults to 0 if not provided. However, it does return an error if
// it exists but cannot be converted to an integer
func getSkip(r *http.Request) (int, error) {
	s := r.URL.Query().Get("skip")
	if s == "" {
		return 0, nil
	}

	sI, err := strconv.Atoi(s)
	if err != nil {
		return 0, errors.Wrap(err, "could not convert skip parameter to int")
	}

	return sI, nil
}

// getTake returns the value of the take query string parameter
// It defaults to 100 if not provided. However, it does return an error if
// it exists but cannot be converted to an integer
func getTake(r *http.Request) (int, error) {
	t := r.URL.Query().Get("take")
	if t == "" {
		return 0, nil
	}

	tI, err := strconv.Atoi(t)
	if err != nil {
		return 0, errors.Wrap(err, "could not convert skip parameter to int")
	}

	return tI, nil
}
