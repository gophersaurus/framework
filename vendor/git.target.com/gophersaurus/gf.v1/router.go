package gf

import (
	"net/http"
	"path/filepath"

	"git.target.com/gophersaurus/gophersaurus/vendor/github.com/codegangsta/negroni"
	"git.target.com/gophersaurus/gophersaurus/vendor/github.com/gorilla/mux"
	render "git.target.com/gophersaurus/gophersaurus/vendor/gopkg.in/unrolled/render.v1"
)

var renderer *render.Render

// Router
type Router struct {
	mux  *mux.Router
	n    *negroni.Negroni
	keys KeyMap
}

func NewRouter(keys KeyMap, indentJson bool) *Router {
	if renderer == nil {
		renderer = render.New(render.Options{
			IndentJSON: indentJson,
		})
	}
	n := negroni.New()

	m := mux.NewRouter()
	n.UseHandler(m)

	r := &Router{mux: m, n: n, keys: keys}
	return r
}

func (r *Router) Middleware(h negroni.Handler) {
	r.n.Use(h)
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.n.ServeHTTP(w, req)
}

func (r *Router) Static(path, dir string) {
	projectPath := filepath.Dir(ConfigPath)
	abs, err := filepath.Abs(filepath.Dir(projectPath + "/" + dir))
	Check(err)
	r.mux.PathPrefix(path).Handler(http.FileServer(http.Dir(abs)))
}

func (r *Router) Get(path string, action Action) {
	route := r.mux.Path(path)
	route.Methods("GET")
	route.Handler(negroni.New(&keyHandler{r.keys}, negroni.Wrap(action)))
}

func (r *Router) Post(path string, action Action) {
	route := r.mux.Path(path)
	route.Methods("POST")
	route.Headers("Content-Type", "application/json")
	route.Handler(negroni.New(&keyHandler{r.keys}, negroni.Wrap(action)))
}

func (r *Router) Put(path string, action Action) {
	route := r.mux.Path(path)
	route.Methods("PUT")
	route.Headers("Content-Type", "application/json")
	route.Handler(negroni.New(&keyHandler{r.keys}, negroni.Wrap(action)))
}

func (r *Router) Patch(path string, action Action) {
	route := r.mux.Path(path)
	route.Methods("PATCH")
	route.Headers("Content-Type", "application/json")
	route.Handler(negroni.New(&keyHandler{r.keys}, negroni.Wrap(action)))
}

func (r *Router) Delete(path string, action Action) {
	route := r.mux.Path(path)
	route.Methods("DELETE")
	route.Handler(negroni.New(&keyHandler{r.keys}, negroni.Wrap(action)))
}

func (r *Router) Resource(path string, c Controller) {

	defaultPathId := path + "/{id}"

	r.Get(path, c.Index)
	r.Post(path, c.Store)
	r.Get(defaultPathId, c.Show)
	r.Delete(defaultPathId, c.Destroy)

	u, ok := c.(Updateable)
	if ok {
		r.Put(defaultPathId, u.Update)
	}

	p, ok := c.(Patchable)
	if ok {
		r.Patch(defaultPathId, p.Apply)
	}
}
