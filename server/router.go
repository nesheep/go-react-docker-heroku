package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/nesheep/go-react-docker-heroku/config"
	"github.com/nesheep/go-react-docker-heroku/frontend"
	"github.com/nesheep/go-react-docker-heroku/handler"
)

func NewRouter(cfg *config.Config) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)

	if cfg.Env == "dev" {
		r.Use(cors.Handler(cors.Options{
			AllowedOrigins:   []string{"*"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Accept", "Content-Type"},
			AllowCredentials: false,
		}))
	}

	h := handler.NewHello()
	f := handler.NewFrontend(frontend.Assets, "build")

	r.Get("/hello/{name}", h.Get)
	r.NotFound(f.ServeHTTP)

	return r
}
