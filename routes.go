package main

import "github.com/macedo/category_service-go/Godeps/_workspace/src/github.com/julienschmidt/httprouter"
import "github.com/macedo/category_service-go/handlers"

type Route struct {
	Name        string
	Method      string
	Path        string
	HandlerFunc httprouter.Handle
}

type Routes []Route

var routes = Routes{
	Route{
		"CategoryIndex",
		"GET",
		"/categories",
		handlers.CategoryIndex,
	},

	Route{
		"CategoryShow",
		"GET",
		"/categories/:categoryId",
		handlers.CategoryShow,
	},
}
