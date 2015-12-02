package router

import (
	"log"
	stdhttp "net/http"
	"os"
	"path"
	"path/filepath"

	"github.com/gophersaurus/gf.v1/http"
	"github.com/gophersaurus/gf.v1/resource"
	"github.com/julienschmidt/httprouter"
)

var endpoints []Endpoint

// Mux describes a HTTP multiplex router.
type Mux struct {
	mux    *httprouter.Router
	mc     MiddlewareChain
	prefix string
}

// Endpoint represents an API URI endpoint.  API endpoints are only used for
// documentation purposes.
type Endpoint struct {
	Type string
	Path string
}

// NewMux returns a new router.
func NewMux() *Mux {

	// create a new HTTP multiplexer
	mux := httprouter.New()
	mux.HandleMethodNotAllowed = false

	// create a new middleware chain
	mc := NewMiddlewareChain()

	// create a new router
	return &Mux{mux: mux, mc: mc}
}

// Endpoints returns the list of registered HTTP endpoints.
func Endpoints() []Endpoint {
	return endpoints
}

// Middleware registers HTTP middlware.
func (m *Mux) Middleware(middleware ...Middleware) Router {
	if len(middleware) > 0 {
		m.mc = m.mc.Append(middleware...)
	}
	return m
}

// FlushMiddleware creates a subrouter with a fresh middleware chain.
func (m *Mux) FlushMiddleware(path string) Router {
	return &Mux{
		mux:    m.mux,           // pass a fresh router
		prefix: m.prefix + path, // pass the path prefix
	}
}

// ServeHTTP satisfies the http.Hander interface. This provides flexiblity and
// compatiblity with the standard http package.
func (m *Mux) ServeHTTP(w stdhttp.ResponseWriter, req *stdhttp.Request) {
	m.mux.ServeHTTP(w, req)
}

// Subrouter creates a new subrouter based on the path prefix and middlweare of its parent.
func (m *Mux) Subrouter(path string) Router {
	return &Mux{
		mux:    m.mux,           // pass a fresh router
		mc:     m.mc,            // pass on the middleware chain
		prefix: m.prefix + path, // pass the path prefix
	}
}

// GET registers a URL path with an Action.
func (m *Mux) GET(uri string, f http.HandlerFunc, middleware ...Middleware) {

	// join and clean router prefix and uri
	url := path.Join(m.prefix, uri)

	// if the path ends with a param use a dynamic formatted action, otherwise
	// statically define the routes for better performance
	if paramEnd(uri) {
		m.mux.GET(url, m.action(f, middleware...))
	} else {
		m.mux.GET(url, m.actionWithFormat("json", f, middleware...))
		m.mux.GET(url+".json", m.actionWithFormat("json", f, middleware...))
		m.mux.GET(url+".xml", m.actionWithFormat("xml", f, middleware...))
		m.mux.GET(url+".yml", m.actionWithFormat("yml", f, middleware...))
	}

	// create an endpoint and append it to the slice of endpoints
	endpoints = append(endpoints, Endpoint{Type: "GET", Path: url})
}

// POST registers a URL path with an Action.
func (r *Mux) POST(uri string, f http.HandlerFunc, m ...Middleware) {

	// join and clean router prefix and uri
	url := path.Join(r.prefix, uri)

	// if the path ends with a param use a dynamic formatted action, otherwise
	// statically define the routes for better performance
	if paramEnd(uri) {
		r.mux.POST(url, r.action(f, m...))
	} else {
		r.mux.POST(url, r.actionWithFormat("json", f, m...))
		r.mux.POST(url+".json", r.actionWithFormat("json", f, m...))
		r.mux.POST(url+".xml", r.actionWithFormat("xml", f, m...))
		r.mux.POST(url+".yml", r.actionWithFormat("yml", f, m...))
	}

	// create an endpoint and append it to the slice of endpoints
	endpoints = append(endpoints, Endpoint{Type: "POST", Path: url})
}

// PATCH registers a URL path with an Action.
func (r *Mux) PATCH(uri string, f http.HandlerFunc, m ...Middleware) {

	// join and clean router prefix and uri
	url := path.Join(r.prefix, uri)

	// if the path ends with a param use a dynamic formatted action, otherwise
	// statically define the routes for better performance
	if paramEnd(uri) {
		r.mux.PATCH(url, r.action(f, m...))
	} else {
		r.mux.PATCH(url, r.actionWithFormat("json", f, m...))
		r.mux.PATCH(url+".json", r.actionWithFormat("json", f, m...))
		r.mux.PATCH(url+".xml", r.actionWithFormat("xml", f, m...))
		r.mux.PATCH(url+".yml", r.actionWithFormat("yml", f, m...))
	}

	// create an endpoint and append it to the slice of endpoints
	endpoints = append(endpoints, Endpoint{Type: "PATCH", Path: url})
}

// PUT registers a URL path with an Action.
func (r *Mux) PUT(uri string, f http.HandlerFunc, m ...Middleware) {

	// join and clean router prefix and uri
	url := path.Join(r.prefix, uri)

	// if the path ends with a param use a dynamic formatted action, otherwise
	// statically define the routes for better performance
	if paramEnd(uri) {
		r.mux.PUT(url, r.action(f, m...))
	} else {
		r.mux.PUT(url, r.actionWithFormat("json", f, m...))
		r.mux.PUT(url+".json", r.actionWithFormat("json", f, m...))
		r.mux.PUT(url+".xml", r.actionWithFormat("xml", f, m...))
		r.mux.PUT(url+".yml", r.actionWithFormat("yml", f, m...))
	}

	// create an endpoint and append it to the slice of endpoints
	endpoints = append(endpoints, Endpoint{Type: "PUT", Path: url})
}

// DELETE registers a URL path with an Action.
func (r *Mux) DELETE(uri string, f http.HandlerFunc, m ...Middleware) {

	// join and clean router prefix and uri
	url := path.Join(r.prefix, uri)

	// if the path ends with a param use a dynamic formatted action, otherwise
	// statically define the routes for better performance
	if paramEnd(uri) {
		r.mux.DELETE(url, r.action(f, m...))
	} else {
		r.mux.DELETE(url, r.actionWithFormat("json", f, m...))
		r.mux.DELETE(url+".json", r.actionWithFormat("json", f, m...))
		r.mux.DELETE(url+".xml", r.actionWithFormat("xml", f, m...))
		r.mux.DELETE(url+".yml", r.actionWithFormat("yml", f, m...))
	}

	// create an endpoint and append it to the slice of endpoints
	endpoints = append(endpoints, Endpoint{Type: "DELETE", Path: url})
}

// action is a private HTTP handler that executes a controller method.
//
// action also takes multiple negroni.Handler objects to create the middleware
// chain for a route.
func (r *Mux) action(h http.Handler, m ...Middleware) httprouter.Handle {

	// return a function that satisfies the httprouter.Handle interface
	return func(w stdhttp.ResponseWriter, request *stdhttp.Request, ps httprouter.Params) {

		// initalize format as an empty string
		format := ""

		// determine format
		if len(ps) > 0 {
			v := ps[len(ps)-1].Value
			l := len(v)
			if l > 4 {
				switch v[l-4:] {
				case ".xml":
					format = "xml"
				case ".yml":
					format = "yml"
				default:
					format = "json"
				}
			}
		}

		resp := http.NewResponse(w, format)
		req := http.NewRequest(request, ps)

		// middleware exists append it
		if len(m) > 0 {
			chain := r.mc.Append(m...).Then(h)
			chain.ServeHTTP(resp, req)
		} else {
			chain := r.mc.Then(h)
			chain.ServeHTTP(resp, req)
		}
	}
}

// paramEnd checks if path ends in a parameter.
func paramEnd(path string) bool {
	for i := len(path) - 2; i >= 0; i-- {
		if path[i] == ':' {
			return true
		}
		if path[i] == '/' {
			return false
		}
	}
	return false
}

// actionWithFormat executes an action method, but also specifies the return format.
func (r *Mux) actionWithFormat(format string, h http.Handler, m ...Middleware) httprouter.Handle {

	// return a function that satisfies the httprouter.Handle interface
	return func(w stdhttp.ResponseWriter, request *stdhttp.Request, ps httprouter.Params) {

		resp := http.NewResponse(w, format)
		req := http.NewRequest(request, ps)

		// middleware exists append it
		if len(m) > 0 {
			chain := r.mc.Append(m...).Then(h)
			chain.ServeHTTP(resp, req)
		} else {
			chain := r.mc.Then(h)
			chain.ServeHTTP(resp, req)
		}
	}
}

// Resource registers a URL path with a Controller that impliments all
// Index, Store, Show, Update, Apply, and Destory Actions.
func (r *Mux) Resource(path, id string, rs resource.Resourcer, m ...Middleware) {

	// concatenate the full url path, including the param id
	pathID := r.prefix + path + "/:" + id

	r.GET(path, rs.Index, m...)
	r.GET(pathID, rs.Show, m...)
	r.POST(path, rs.Store, m...)
	r.PATCH(pathID, rs.Apply, m...)
	r.PUT(pathID, rs.Update, m...)
	r.DELETE(pathID, rs.Destroy, m...)
}

// Static registers a URL path with a public directory to serve its content.
// This directory is meant to serve public static files such as image files,
// CSS files, JavaScript files, and more.
func (r *Mux) Static(uri, dir string) {

	// clean root paths
	uri = path.Clean(uri)
	if uri == "/" {
		uri = ""
	}

	// join and clean router prefix and uri
	url := path.Join(r.prefix, uri)
	files := path.Join(url, "*filepath")

	// directory path is valid
	if _, err := os.Stat(dir); err != nil {

		// since dir is not a valid directory path, assume the path given is
		// relative to the binary executing
		current, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			log.Fatalln(err)
		}

		// join file paths
		dir = filepath.Join(current, dir)

		// serve all files in the directory
		r.mux.ServeFiles(files, stdhttp.Dir(dir))
		return
	}

	// serve all files in the directory
	r.mux.ServeFiles(files, stdhttp.Dir(dir))
}
