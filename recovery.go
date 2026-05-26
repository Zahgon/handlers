package handlers

import (
	"net/http"
)

// RecoveryHandlerLogger is an interface used by the recovering handler to print logs.
type RecoveryHandlerLogger interface {
	Println(...interface{})
}

type recoveryHandler struct {
	handler    http.Handler
	logger     RecoveryHandlerLogger
	printStack bool
}

// RecoveryOption provides a functional approach to define
// configuration for a handler; such as setting the logging
// whether or not to print stack traces on panic.
type RecoveryOption func(http.Handler)

func parseRecoveryOptions(h http.Handler, opts ...RecoveryOption) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

// RecoveryHandler is HTTP middleware that recovers from a panic,
// logs the panic, writes http.StatusInternalServerError, and
// continues to the next handler.
//
// Example:
//
//	r := mux.NewRouter()
//	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
//		panic("Unexpected error!")
//	})
//
//	http.ListenAndServe(":1123", handlers.RecoveryHandler()(r))
func RecoveryHandler(opts ...RecoveryOption) func(h http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return nil
}

// RecoveryLogger is a functional option to override
// the default logger.
func RecoveryLogger(logger RecoveryHandlerLogger) RecoveryOption {
	_ = "STUB: not implemented"
	return *new(RecoveryOption)
}

//nolint:errcheck //TODO:
// @bharat-rajani should return type-assertion error but would break the API?

// PrintRecoveryStack is a functional option to enable
// or disable printing stack traces on panic.
func PrintRecoveryStack(shouldPrint bool) RecoveryOption {
	_ = "STUB: not implemented"
	return *new(RecoveryOption)
}

//nolint:errcheck //TODO:
// @bharat-rajani should return type-assertion error but would break the API?

func (h recoveryHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (h recoveryHandler) log(v ...interface{}) { _ = "STUB: not implemented"; return }
