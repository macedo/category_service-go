package api

import (
	"github.com/macedo/category_service-go/Godeps/_workspace/src/github.com/gorilla/mux"
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Path        string
	HandlerFunc http.Handler
}

type Routes []Route

func NewRouter() *mux.Router {
	return mux.NewRouter().StrictSlash(true)
}
