package rhrouter

import (
	"errors"
	"fmt"
	"net/http"
)

type Router struct {
	tree *Tree
}

type route struct {
	methods []string
	path    string
	handler http.Handler
}

var (
	tmpRoute = &route{}
	// Error for not found.
	ErrNotFound = errors.New("no matching route was found")
	// Error for method not allowed.
	ErrMethodNotAllowed = errors.New("methods is not allowed")
)

func NewRouter() *Router {
	return &Router{
		tree: NewTree(),
	}
}

func (r *Router) Methods(methods ...string) *Router {
	tmpRoute.methods = append(tmpRoute.methods, methods...)
	return r
}

// Handler sets a handler
func (r *Router) Handler(path string, handler http.Handler) {
	tmpRoute.handler = handler
	tmpRoute.path = path
	r.Handle()
}

// Handle insertes handlers into tree
func (r *Router) Handle() {
	err := r.tree.Insert(tmpRoute.methods, tmpRoute.path, tmpRoute.handler)
	if err != nil {
		fmt.Println(err)
	}
	tmpRoute = &route{}
}

//	ServeHTTP dispatches the request to the handler whose
//
// pattern most closely matches the request URL.
//
//	'https://cs.opensource.google/go/go/+/refs/tags/go1.17.2:src/net/http/server.go;l=2045'
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	method := req.Method
	path := req.URL.Path
	result, err := r.tree.Search(method, path)
	if err != nil {
		status := handleErr(err)
		w.WriteHeader(status)
		return
	}
	h := result.actions.handler
	h.ServeHTTP(w, req)
}

func handleErr(err error) int {
	var status int
	switch err {
	case ErrMethodNotAllowed:
		status = http.StatusMethodNotAllowed
	case ErrNotFound:
		status = http.StatusNotFound
	}
	return status
}
