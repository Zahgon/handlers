package handlers

import (
	"net/http"
)

// CORSOption represents a functional option for configuring the CORS middleware.
type CORSOption func(*cors) error

type cors struct {
	h                      http.Handler
	allowedHeaders         []string
	allowedMethods         []string
	allowedOrigins         []string
	allowedOriginValidator OriginValidator
	exposedHeaders         []string
	maxAge                 int
	ignoreOptions          bool
	allowCredentials       bool
	optionStatusCode       int
}

// OriginValidator takes an origin string and returns whether or not that origin is allowed.
type OriginValidator func(string) bool

var (
	defaultCorsOptionStatusCode = http.StatusOK
	defaultCorsMethods          = []string{http.MethodGet, http.MethodHead, http.MethodPost}
	defaultCorsHeaders          = []string{"Accept", "Accept-Language", "Content-Language", "Origin"}
	// (WebKit/Safari v9 sends the Origin header by default in AJAX requests).
)

const (
	corsOptionMethod           string = http.MethodOptions
	corsAllowOriginHeader      string = "Access-Control-Allow-Origin"
	corsExposeHeadersHeader    string = "Access-Control-Expose-Headers"
	corsMaxAgeHeader           string = "Access-Control-Max-Age"
	corsAllowMethodsHeader     string = "Access-Control-Allow-Methods"
	corsAllowHeadersHeader     string = "Access-Control-Allow-Headers"
	corsAllowCredentialsHeader string = "Access-Control-Allow-Credentials"
	corsRequestMethodHeader    string = "Access-Control-Request-Method"
	corsRequestHeadersHeader   string = "Access-Control-Request-Headers"
	corsOriginHeader           string = "Origin"
	corsVaryHeader             string = "Vary"
	corsOriginMatchAll         string = "*"
)

func (ch *cors) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// A configuration of * is different than explicitly setting an allowed
// origin. Returning arbitrary origin headers in an access control allow
// origin header is unsafe and is not required by any use case.

// CORS provides Cross-Origin Resource Sharing middleware.
// Example:
//
//	import (
//	    "net/http"
//
//	    "github.com/gorilla/handlers"
//	    "github.com/gorilla/mux"
//	)
//
//	func main() {
//	    r := mux.NewRouter()
//	    r.HandleFunc("/users", UserEndpoint)
//	    r.HandleFunc("/projects", ProjectEndpoint)
//
//	    // Apply the CORS middleware to our top-level router, with the defaults.
//	    http.ListenAndServe(":8000", handlers.CORS()(r))
//	}
func CORS(opts ...CORSOption) func(http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return nil
}

func parseCORSOptions(opts ...CORSOption) *cors { _ = "STUB: not implemented"; return nil }

//TODO: @bharat-rajani, return error to caller if not nil?

//
// Functional options for configuring CORS.
//

// AllowedHeaders adds the provided headers to the list of allowed headers in a
// CORS request.
// This is an append operation so the headers Accept, Accept-Language,
// and Content-Language are always allowed.
// Content-Type must be explicitly declared if accepting Content-Types other than
// application/x-www-form-urlencoded, multipart/form-data, or text/plain.
func AllowedHeaders(headers []string) CORSOption {
	_ = "STUB: not implemented"
	return *new(CORSOption)
}

// AllowedMethods can be used to explicitly allow methods in the
// Access-Control-Allow-Methods header.
// This is a replacement operation so you must also
// pass GET, HEAD, and POST if you wish to support those methods.
func AllowedMethods(methods []string) CORSOption {
	_ = "STUB: not implemented"
	return *new(CORSOption)
}

// AllowedOrigins sets the allowed origins for CORS requests, as used in the
// 'Allow-Access-Control-Origin' HTTP header.
// Note: Passing in a []string{"*"} will allow any domain.
func AllowedOrigins(origins []string) CORSOption {
	_ = "STUB: not implemented"
	return *new(CORSOption)
}

// AllowedOriginValidator sets a function for evaluating allowed origins in CORS requests, represented by the
// 'Allow-Access-Control-Origin' HTTP header.
func AllowedOriginValidator(fn OriginValidator) CORSOption {
	_ = "STUB: not implemented"
	return *new(CORSOption)
}

// OptionStatusCode sets a custom status code on the OPTIONS requests.
// Default behaviour sets it to 200 to reflect best practices. This is option is not mandatory
// and can be used if you need a custom status code (i.e 204).
//
// More informations on the spec:
// https://fetch.spec.whatwg.org/#cors-preflight-fetch
func OptionStatusCode(code int) CORSOption { _ = "STUB: not implemented"; return *new(CORSOption) }

// ExposedHeaders can be used to specify headers that are available
// and will not be stripped out by the user-agent.
func ExposedHeaders(headers []string) CORSOption {
	_ = "STUB: not implemented"
	return *new(CORSOption)
}

// MaxAge determines the maximum age (in seconds) between preflight requests. A
// maximum of 10 minutes is allowed. An age above this value will default to 10
// minutes.
func MaxAge(age int) CORSOption { _ = "STUB: not implemented"; return *new(CORSOption) }

// Maximum of 10 minutes.

// IgnoreOptions causes the CORS middleware to ignore OPTIONS requests, instead
// passing them through to the next handler. This is useful when your application
// or framework has a pre-existing mechanism for responding to OPTIONS requests.
func IgnoreOptions() CORSOption { _ = "STUB: not implemented"; return *new(CORSOption) }

// AllowCredentials can be used to specify that the user agent may pass
// authentication details along with the request.
func AllowCredentials() CORSOption { _ = "STUB: not implemented"; return *new(CORSOption) }

func (ch *cors) isOriginAllowed(origin string) bool { _ = "STUB: not implemented"; return false }

func (ch *cors) isMatch(needle string, haystack []string) bool {
	_ = "STUB: not implemented"
	return false
}
