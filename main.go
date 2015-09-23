package main

import (
	"log"
	"net/http"

	"github.com/macedo/category_service-go/Godeps/_workspace/src/gopkg.in/caarlos0/env.v1"
)

var cfg config

type config struct {
	Port        string `env:"PORT" envDefault:"1987"`
	DatabaseURL string `env:"DATABASE_URL" envDefault:"postgres://category_service-go:@127.0.0.1:5432/category_service-go?sslmode=disable"`
}

func init() {
	cfg = config{}
	env.Parse(&cfg)
}

func main() {
	router := NewRouter()
	log.Fatal(http.ListenAndServe(":"+cfg.Port, router))
}
