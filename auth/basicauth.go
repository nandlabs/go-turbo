package auth

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"
)

type (
	BasicAuthFilter struct {
		Validator BasicAuthValidator
	}

	// BasicAuthValidator expects username and password
	BasicAuthValidator func(string, string) (bool, error)
)

var (
	DefaultBasicAuthFilterConfig = BasicAuthFilter{}
)

func (ba *BasicAuthFilter) Apply(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Basic Auth Implementation
		if ba.Validator == nil {
			panic("basic-auth filter requires a validator function")
		}
		// perform pre-requisite checks
		auth := r.Header.Get(HeaderAuthorization)
		l := len(basic)
		if len(auth) > l+1 && strings.EqualFold(auth[:l], basic) {
			basicAuth, err := base64.StdEncoding.DecodeString(auth[l+1:])
			if err != nil {
				fmt.Print("error decoding basic-auth token")
				w.WriteHeader(http.StatusInternalServerError)
				//panic("unable to decode basic-auth token")
			}
			fmt.Sprintf("basic token: %s", basicAuth)
			tokenUsername := strings.Split(string(basicAuth), ":")[0]
			tokenPassword := strings.Split(string(basicAuth), ":")[1]

			valid, err := ba.Validator(tokenUsername, tokenPassword)
			if err != nil {
				fmt.Print("error validating basic-auth token")
				w.WriteHeader(http.StatusInternalServerError)
				//panic("error validating token")
			} else if valid {
				next.ServeHTTP(w, r)
			}
		}
		// handle in case of authorization token not sent
		w.WriteHeader(http.StatusUnauthorized)
	})
}

func CreateBasicAuthAuthenticator(fn BasicAuthValidator) BasicAuthFilter {
	filterConfig := DefaultBasicAuthFilterConfig
	filterConfig.Validator = fn
	return filterConfig
}
