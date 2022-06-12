package turbo

import (
	"net/http"
	"path"
)

type (
	HttpError struct {
		statusCode int
		message    string
	}
)

//Common constants used throughout
const (
	PathSeparator       = "/"
	GET                 = "GET"
	HEAD                = "HEAD"
	POST                = "POST"
	PUT                 = "PUT"
	DELETE              = "DELETE"
	OPTIONS             = "OPTIONS"
	TRACE               = "TRACE"
	PATCH               = "PATCH"
	Basic               = "basic"
	HeaderAuthorization = "Authorization"
)

var Methods = map[string]string{
	GET:     GET,
	HEAD:    HEAD,
	POST:    POST,
	PUT:     PUT,
	DELETE:  DELETE,
	OPTIONS: OPTIONS,
	TRACE:   TRACE,
	PATCH:   PATCH,
}

//refinePath Borrowed from the golang's net/turbo package
func refinePath(p string) string {
	if p == "" {
		return "/"
	}
	if p[0] != '/' {
		p = "/" + p
	}
	rp := path.Clean(p)
	if p[len(p)-1] == '/' && rp != "/" {
		rp += "/"
	}
	return rp
}

//endpointNotFound to check for the request endpoint
func endpointNotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	_, err := w.Write([]byte("Endpoint Not Found : " + r.URL.Path + "\n"))
	if err != nil {
		return
	}
}

//endpointNotFoundHandler when a requested endpoint is not found in the registered route's this handler is invoked
func endpointNotFoundHandler() http.Handler {
	return http.HandlerFunc(endpointNotFound)
}

//methodNotAllowed to check for the supported method for the incoming request
func methodNotAllowed(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusMethodNotAllowed)
	_, err := w.Write([]byte("Requested Method : " + r.Method + " not supported for Endpoint : " + r.URL.Path + "\n"))
	if err != nil {
		return
	}
}

//methodNotAllowedHandler when a requested method is not allowed in the registered route's method list this handler is invoked
func methodNotAllowedHandler() http.Handler {
	return http.HandlerFunc(methodNotAllowed)
}

func (httpError *HttpError) errorMethod(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(httpError.statusCode)
	_, err := w.Write([]byte(httpError.message))
	if err != nil {
		return
	}
}

func UnAuthorizedHandler() http.Handler {
	httpError := &HttpError{
		statusCode: http.StatusUnauthorized,
		message:    "Incoming request cannot be authorized \n",
	}
	return http.HandlerFunc(httpError.errorMethod)
}

func InvalidTokenHandler() http.Handler {
	httpError := &HttpError{
		statusCode: http.StatusInternalServerError,
		message:    "Invalid Token provided for the request \n",
	}
	return http.HandlerFunc(httpError.errorMethod)
}

func ErrorDecodingTokenHandler() http.Handler {
	httpError := &HttpError{
		statusCode: http.StatusBadRequest,
		message:    "Error decoding authorization token \n",
	}
	return http.HandlerFunc(httpError.errorMethod)
}
