package middleware

import (
	"errors"
	"net"

	"github.com/gophersaurus/gf.v1/http"
)

// KeyMiddleware checks if keys are valid.
type Keys struct {
	success http.Handler
	Keys    map[string][]string
}

// NewKeyMiddleware returns a KeyHandler.
func NewKeys(keys map[string][]string) Keys {

	for _, whitelist := range keys {
		for i, uri := range whitelist {
			if uri == "localhost" {
				whitelist[i] = "::1"
			}
		}
	}

	return Keys{Keys: keys}
}

// Do takes a handler and executes key middleware.
func (k Keys) Do(h http.Handler) http.Handler {
	k.success = h
	return k
}

// Check takes a key and checks it.
func (k Keys) Check(key, remoteAddr string) error {

	// key exists in the map
	if whitelist, ok := k.Keys[key]; ok {

		// no whitelist is provided, so any request is ok
		if len(whitelist) == 0 {
			return nil
		}

		// range over each whitelisted URI
		for _, uri := range whitelist {
			if uri == "all" || uri == "*" || uri == remoteAddr {
				return nil
			}
		}
	}

	return errors.New(http.InvalidPermission)
}

// ServeHTTP fulfills the http package interface for middlewares.
func (k Keys) ServeHTTP(resp http.Responder, req *http.Request) {

	// get client IP address
	ip, _, _ := net.SplitHostPort(req.RemoteAddr)

	// sneaky little proxies and loadbalancers, wicked, tricksy, false!
	if proxy := req.Header.Get("X-FORWARDED-FOR"); len(proxy) > 0 {
		ip = proxy
	}

	// API key parameter in URL
	if key := req.URL.Query().Get("key"); len(key) > 0 {
		if err := k.Check(key, ip); err != nil {
			resp.WriteErrs(req, err.Error())
			return
		}
		k.success.ServeHTTP(resp, req)
		return
	}

	// API key in header
	if key := req.Header.Get("API-Key"); len(key) > 0 {
		if err := k.Check(key, ip); err != nil {
			resp.WriteErrs(req, err.Error())
			return
		}
		k.success.ServeHTTP(resp, req)
		return
	}

	// no API key provided
	resp.WriteErrs(req, http.InvalidPermission)
}
