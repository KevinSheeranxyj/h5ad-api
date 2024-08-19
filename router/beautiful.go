package router

import (
	"net/http"
)

type middleware func(http.HandlerFunc) http.HandlerFunc

type Router struct {
	//
	middlewareChain []middleware
	//
	mux map[string]http.HandlerFunc
}

func NewRouter() *Router {
	return &Router{middlewareChain: nil, mux: make(map[string]http.HandlerFunc)}
}

func (r *Router) Use(m middleware) {
	r.middlewareChain = append(r.middlewareChain, m)
}

func (r *Router) Add(route string, h http.HandlerFunc) {
	var MergeHandler = h
	for i := len(r.middlewareChain) - 1; i >= 0; i-- {
		MergeHandler = r.middlewareChain[i](MergeHandler)
	}
	r.mux[route] = MergeHandler
}

func (r *Router) Load() {
	for router, hand := range r.mux {
		http.Handle(router, hand)
	}
}
