package gf

import (
	"errors"
	"net/http"
)

type keyHandler struct {
	keys KeyMap
}

func (k *keyHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	queryKey := r.URL.Query().Get("key")
	headerKey := r.Header.Get("API-Key")
	if !k.isKeyValid(r.RemoteAddr, queryKey, headerKey) {
		buildResponse(rw).RespondWithErr(errors.New(InvalidPermissions))
		return
	}
	next(rw, r)
}

func (k *keyHandler) isKeyValid(sender string, keys ...string) bool {
	if len(keys) <= 0 || keys == nil {
		return true
	}
	var conf *KeyConfig
	i := 0
	for conf == nil && i < len(keys) {
		conf = k.keys.Get(keys[i])
		i++
	}
	if conf == nil {
		return false
	}
	if !conf.Status {
		return false
	}
	if len(conf.Urls) > 0 {
		for _, url := range conf.Urls {
			if url == sender {
				return true
			}
		}
		return false
	}
	return true
}
