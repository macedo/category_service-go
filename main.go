package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/macedo/category_service-go/api"
)

func main() {
	api.NewServer()
	api.StartServer()
	api.InitApi()

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-c

	api.StopServer()
}
