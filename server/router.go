package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/nesheep/go-react-docker-heroku/handler"
)

func NewRouter() http.Handler {
	r := chi.NewRouter()

	hh := handler.NewHelloHander()
	r.Get("/hello/{name}", hh.Hello)

	return r
}
