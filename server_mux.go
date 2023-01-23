package main

import (
	"fmt"
	"net/http"
	rhrouter "v1/rh_router"
)

func indexHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "retuning GET /")
	})
}

func fooHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			fmt.Fprintf(w, "retuning GET /foo")
		case http.MethodPost:
			fmt.Fprintf(w, " retuning POST /foo")
		default:
			fmt.Fprintf(w, " retuning Not Found")
		}
	})
}

func fooBarHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, " retuning GET /foo/bar")
	})
}

func fooBarBazHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, " retuning GET /foo/bar/baz")
	})
}

func barHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, " retuning GET /bar")
	})
}

func bazHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, " retuning GET /baz")
	})
}

//	'ServeMux' is an HTTP request multiplexer
//
// This multiplexer is responsible for checking the URL of the request against the registered patterns and calling the most matching handler (the function that returns the response
func createMuxServer() {
	fmt.Println("Hello go_router!")
	r := rhrouter.NewRouter()

	r.Methods(http.MethodGet).Handler(`/`, indexHandler())
	r.Methods(http.MethodGet, http.MethodPost).Handler(`/foo`, fooHandler())
	r.Methods(http.MethodGet).Handler(`/foo/bar`, fooBarHandler())
	r.Methods(http.MethodGet).Handler(`/foo/bar/baz`, fooBarBazHandler())
	r.Methods(http.MethodGet).Handler(`/bar`, barHandler())
	r.Methods(http.MethodGet).Handler(`/baz`, bazHandler())

	http.ListenAndServe(":9000", r)
}
