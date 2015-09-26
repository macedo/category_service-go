package api

import (
	l4g "github.com/macedo/category_service-go/Godeps/_workspace/src/code.google.com/p/log4go"
	"github.com/macedo/category_service-go/Godeps/_workspace/src/github.com/braintree/manners"
	"github.com/macedo/category_service-go/Godeps/_workspace/src/github.com/gorilla/mux"
	"net/http"
	"os"
	"time"
)

type Server struct {
	Server *manners.GracefulServer
	Router *mux.Router
}

var Srv *Server

func NewServer() {
	l4g.Info("Server is initializing...")

	router := NewRouter()

	Srv = &Server{
		Router: router,
		Server: manners.NewWithServer(&http.Server{
			Addr:           ":" + os.Getenv("PORT"),
			Handler:        router,
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10. * time.Second,
			MaxHeaderBytes: 1 << 20,
		}),
	}
}

func StartServer() {
	l4g.Info("Starting Server...")
	l4g.Info("Server is listening on " + os.Getenv("PORT"))

	go func() {
		err := Srv.Server.ListenAndServe()

		if err != nil {
			l4g.Critical("Error starting server, err:%v", err)
			time.Sleep(time.Second)
			panic("Error starting server " + err.Error())
		}
	}()
}

func StopServer() {
	l4g.Info("Stopping Server...")
	Srv.Server.Close()
	l4g.Info("Server stopped")
}
