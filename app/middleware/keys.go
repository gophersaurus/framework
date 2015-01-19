package middleware

import (
	"net/http"

	"git.target.com/gophersaurus/gf.v1"
)

// KeyHandler contains the KeyMap.
type KeyHandler struct {
	keys gf.KeyMap
}

func NewKeyHandler(keys gf.KeyMap) *KeyHandler {
	return &KeyHandler{keys}
}

// ServeHTTP fufills the http package interface for HTTP middleware.
func (k *KeyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	queryKey := r.URL.Query().Get("key")
	headerKey := r.Header.Get("API-Key")
	if !k.isKeyValid(r.RemoteAddr, queryKey, headerKey) {
		resp := gf.NewResponse(w)
		resp.RespondWithErr(gf.InvalidPermission)
		return
	}
	next(w, r)
}

// isKeyValid checks if a key is valid.
func (k *KeyHandler) isKeyValid(sender string, keys ...string) bool {
	if len(keys) <= 0 || keys == nil {
		return true
	}
	var conf *gf.KeyConfig
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
	if len(conf.Whitelist) > 0 {
		for _, url := range conf.Whitelist {
			if url == sender {
				return true
			}
		}
		return false
	}
	return true
}
