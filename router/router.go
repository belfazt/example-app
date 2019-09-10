package router

import (
	"net/http"
)

// Register mounts our HTTP handler on the mux.
//
// Register is a typical top-level application function: it takes a generic
// type like ServeMux, which typically comes from a third-party library, and
// introduces it to a type that contains our application logic. In this case,
// that introduction consists of registering an HTTP handler. Other typical
// examples include registering RPC procedures and starting queue consumers.
//
// Fx calls these functions invocations, and they're treated differently from
// the constructor functions above. Their arguments are still supplied via
// dependency injection and they may still return an error to indicate
// failure, but any other return values are ignored.
//
// Unlike constructors, invocations are called eagerly. See the main function
// below for details.
func Register(mux *http.ServeMux, h http.Handler) {
	mux.Handle("/", h)
}
