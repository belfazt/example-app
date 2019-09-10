package main

import (
	handler "github.com/belfazt/example-app/handler"
	logger "github.com/belfazt/example-app/logger"
	mux "github.com/belfazt/example-app/mux"
	router "github.com/belfazt/example-app/router"

	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		// Provide all the constructors we need, which teaches Fx how we'd like to
		// construct the *log.Logger, http.Handler, and *http.ServeMux types.
		// Remember that constructors are called lazily, so this block doesn't do
		// much on its own.
		fx.Provide(
			logger.NewLogger,
			handler.NewHandler,
			mux.NewMux,
		),
		// Since constructors are called lazily, we need some invocations to
		// kick-start our application. In this case, we'll use Register. Since it
		// depends on an http.Handler and *http.ServeMux, calling it requires Fx
		// to build those types using the constructors above. Since we call
		// NewMux, we also register Lifecycle hooks to start and stop an HTTP
		// server.
		fx.Invoke(router.Register),
	)

	app.Run()
}
