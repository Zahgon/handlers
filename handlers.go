// Copyright 2013 The Gorilla Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package handlers

import (
	"bufio"
	"net"
	"net/http"
)

// MethodHandler is an http.Handler that dispatches to a handler whose key in the
// MethodHandler's map matches the name of the HTTP request's method, eg: GET
//
// If the request's method is OPTIONS and OPTIONS is not a key in the map then
// the handler responds with a status of 200 and sets the Allow header to a
// comma-separated list of available methods.
//
// If the request's method doesn't match any of its keys the handler responds
// with a status of HTTP 405 "Method Not Allowed" and sets the Allow header to a
// comma-separated list of available methods.
type MethodHandler map[string]http.Handler

func (h MethodHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	_ = "STUB: not implemented"
	return
}

// responseLogger is wrapper of http.ResponseWriter that keeps track of its HTTP
// status code and body size.
type responseLogger struct {
	w      http.ResponseWriter
	status int
	size   int
}

func (l *responseLogger) Write(b []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (l *responseLogger) WriteHeader(s int) { _ = "STUB: not implemented"; return }

func (l *responseLogger) Status() int { _ = "STUB: not implemented"; return 0 }

func (l *responseLogger) Size() int { _ = "STUB: not implemented"; return 0 }

func (l *responseLogger) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil, nil
}

// The status will be StatusSwitchingProtocols if there was no error and
// WriteHeader has not been called yet

// isContentType validates the Content-Type header matches the supplied
// contentType. That is, its type and subtype match.
func isContentType(h http.Header, contentType string) bool { _ = "STUB: not implemented"; return false }

// ContentTypeHandler wraps and returns a http.Handler, validating the request
// content type is compatible with the contentTypes list. It writes a HTTP 415
// error if that fails.
//
// Only PUT, POST, and PATCH requests are considered.
func ContentTypeHandler(h http.Handler, contentTypes ...string) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

const (
	// HTTPMethodOverrideHeader is a commonly used
	// http header to override a request method.
	HTTPMethodOverrideHeader = "X-HTTP-Method-Override"
	// HTTPMethodOverrideFormKey is a commonly used
	// HTML form key to override a request method.
	HTTPMethodOverrideFormKey = "_method"
)

// HTTPMethodOverrideHandler wraps and returns a http.Handler which checks for
// the X-HTTP-Method-Override header or the _method form key, and overrides (if
// valid) request.Method with its value.
//
// This is especially useful for HTTP clients that don't support many http verbs.
// It isn't secure to override e.g a GET to a POST, so only POST requests are
// considered.  Likewise, the override method can only be a "write" method: PUT,
// PATCH or DELETE.
//
// Form method takes precedence over header method.
func HTTPMethodOverrideHandler(h http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}
