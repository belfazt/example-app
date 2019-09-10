package handler

import (
	"log"
	"net/http"
)

// NewHandler constructs a simple HTTP handler. Since it returns an
// http.Handler, Fx will treat NewHandler as the constructor for the
// http.Handler type.
//
// Like many Go functions, NewHandler also returns an error. If the error is
// non-nil, Go convention tells the caller to assume that NewHandler failed
// and the other returned values aren't safe to use. Fx understands this
// idiom, and assumes that any function whose last return value is an error
// follows this convention.
//
// Unlike NewLogger, NewHandler has formal parameters. Fx will interpret these
// parameters as dependencies: in order to construct an HTTP handler,
// NewHandler needs a logger. If the application has access to a *log.Logger
// constructor (like NewLogger above), it will use that constructor or its
// cached output and supply a logger to NewHandler. If the application doesn't
// know how to construct a logger and needs an HTTP handler, it will fail to
// start.
//
// Functions may also return multiple objects. For example, we could combine
// NewHandler and NewLogger into a single function:
//
//   func NewHandlerAndLogger() (*log.Logger, http.Handler, error)
//
// Fx also understands this idiom, and would treat NewHandlerAndLogger as the
// constructor for both the *log.Logger and http.Handler types. Just like
// constructors for a single type, NewHandlerAndLogger would be called at most
// once, and both the handler and the logger would be cached and reused as
// necessary.
func NewHandler(logger *log.Logger) (http.Handler, error) {
	logger.Print("Executing NewHandler.")
	return http.HandlerFunc(func(http.ResponseWriter, *http.Request) {
		logger.Print("Got a request.")
	}), nil
}
