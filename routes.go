package main

import "github.com/macedo/category_service-go/Godeps/_workspace/src/github.com/julienschmidt/httprouter"

type Route struct {
	Name        string
	Method      string
	Path        string
	HandlerFunc httprouter.Handle
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},

	Route{
		"CategoryIndex",
		"GET",
		"/categories",
		CategoryIndex,
	},

	Route{
		"CategoryShow",
		"GET",
		"/categories/:categoryId",
		CategoryShow,
	},
}
