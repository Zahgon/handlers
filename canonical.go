package handlers

import (
	"net/http"
)

type canonical struct {
	h      http.Handler
	domain string
	code   int
}

// CanonicalHost is HTTP middleware that re-directs requests to the canonical
// domain. It accepts a domain and a status code (e.g. 301 or 302) and
// re-directs clients to this domain. The existing request path is maintained.
//
// Note: If the provided domain is considered invalid by url.Parse or otherwise
// returns an empty scheme or host, clients are not re-directed.
//
// Example:
//
//	r := mux.NewRouter()
//	canonical := handlers.CanonicalHost("http://www.gorillatoolkit.org", 302)
//	r.HandleFunc("/route", YourHandler)
//
//	log.Fatal(http.ListenAndServe(":7000", canonical(r)))
func CanonicalHost(domain string, code int) func(h http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return nil
}

func (c canonical) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// Call the next handler if the provided domain fails to parse.

// Call the next handler if the scheme or host are empty.
// Note that url.Parse won't fail on in this case.

// Re-build the destination URL

// cleanHost cleans invalid Host headers by stripping anything after '/' or ' '.
// This is backported from Go 1.5 (in response to issue #11206) and attempts to
// mitigate malformed Host headers that do not match the format in RFC7230.
func cleanHost(in string) string { _ = "STUB: not implemented"; return "" }
