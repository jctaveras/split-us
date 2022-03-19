package router

import "net/http"

type router struct {
	mux *http.ServeMux
}

func newRouter() *router {
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

var Routes = newRouter()
