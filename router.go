package main

import "github.com/macedo/category_service-go/Godeps/_workspace/src/github.com/julienschmidt/httprouter"

func NewRouter() *httprouter.Router {
	router := httprouter.New()

	for _, route := range routes {
		router.Handle(route.Method, route.Path, route.HandlerFunc)
	}

	return router
}
