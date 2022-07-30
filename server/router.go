package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/nesheep/go-react-docker-heroku/config"
	"github.com/nesheep/go-react-docker-heroku/handler"
)

func NewRouter(cfg *config.Config, html http.FileSystem) http.Handler {
	r := chi.NewRouter()

	if cfg.Env == "dev" {
		r.Use(cors.Handler(cors.Options{
			AllowedOrigins:   []string{"*"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Accept", "Content-Type"},
			AllowCredentials: false,
		}))
	}

	hh := handler.NewHelloHander()
	r.Get("/*", http.FileServer(html).ServeHTTP)
	r.Get("/hello/{name}", hh.Hello)

	return r
}
