package api

import (
	"fmt"
	l4g "github.com/macedo/category_service-go/Godeps/_workspace/src/code.google.com/p/log4go"
	"net/http"
)

var routes = Routes{
	Route{
		"CategoryIndex",
		"GET",
		"/categories",
		ApiHandler(categoryIndex),
	},
}

func InitCategory() {
	l4g.Debug("Initializing categories api routes")

	for _, route := range routes {
		Srv.Router.
			Methods(route.Method).
			Path(route.Path).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
}

func categoryIndex(c *Controller, w http.ResponseWriter, r *http.Request) {
	fmt.Println("adadasdad")
}
