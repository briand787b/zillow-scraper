package controller

import (
	"context"
	"fmt"
	"log"
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
)

// Serve is a blocking function that serves HTTP
func Serve(port int, l plog.Logger, ps model.PropertyStore) {
	// pc := NewPropertyController(l, ps)
	mw, err := NewMiddleware(l, uuid.New(), os.Getenv(CorsEnvVarKey))
	if err != nil {
		log.Fatalln(err)
	}

	// initialize router
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(render.SetContentType(render.ContentTypeJSON))
	r.Use(mw.spanAndTrace)
	r.Use(mw.logRoute)
	r.Use(mw.disableCORS)

	r.Route("/properties", func(r chi.Router) {
		// r.Get("/", mc.HandleGetAllRoot)
		// r.Post("/", mc.HandleCreate)
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
		log.Fatalln("could not print API routes: ", err)
	}

	log.Fatalln(http.ListenAndServe(fmt.Sprintf(":%v", port), r))

}
