package controller

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"reflect"
	"runtime"

	"zcrapr/core/model"
	"zcrapr/core/plog"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

// Serve is a blocking function that serves HTTP
func Serve(port int, l plog.Logger, cs model.CaptureStore, ps model.PropertyStore) error {
	cc := NewCaptureController(l, cs, ps)
	pc := NewPropertyController(l, ps)
	mw, err := NewMiddleware(l, uuid.New(), os.Getenv(CorsEnvVarKey))
	if err != nil {
		return errors.Wrap(err, "could not create new middleware")
	}

	// initialize router
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(render.SetContentType(render.ContentTypeJSON))
	r.Use(mw.spanAndTrace)
	r.Use(mw.logRoute)
	r.Use(mw.disableCORS)

	r.Route("/properties", func(r chi.Router) {
		r.Post("/", pc.HandleCreate)
		r.With(mw.skipTake).Get("/", pc.HandleGetAll)
		r.Route("/{property_id}", func(r chi.Router) {
			r.With(pc.propertyCtx).Get("/", pc.HandleGetByID)
		})
	})

	r.Route("/captures", func(r chi.Router) {
		r.Post("/", cc.HandleCreate)
	})

	ctx := plog.StoreSpanIDTraceID(context.Background(), "main", "main")
	walkFn := func(method, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		funcName := runtime.FuncForPC(reflect.ValueOf(handler).Pointer()).Name()
		l.Info(ctx, "API Route",
			"method", method,
			"route", route,
			"handler", funcName,
		)
		return nil
	}

	if err := chi.Walk(r, walkFn); err != nil {
		return errors.Wrap(err, "could not print API routes")
	}

	if err := http.ListenAndServe(fmt.Sprintf(":%v", port), r); err != nil {
		return errors.Wrap(err, "could not listen and serve")
	}

	return nil
}
