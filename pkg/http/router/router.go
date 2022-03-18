package router

import "net/http"

type Router interface {
	Handlers() *http.ServeMux
	GET(string, http.HandlerFunc)
	POST(string, http.HandlerFunc)
}

type router struct {
	mux *http.ServeMux
}

func NewRouter() Router {
	mux := http.NewServeMux()
	return &router{mux}
}

func (r *router) Handlers() *http.ServeMux {
	return r.mux
}

func (r *router) GET(pattern string, fn http.HandlerFunc)  {
	r.mux.HandleFunc(pattern, fn)
}

func (r *router) POST(pattern string, fn http.HandlerFunc) {
	r.mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		fn(w, r)
	})
}
