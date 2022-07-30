package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/nesheep/go-react-docker-heroku/handler"
)

func NewRouter(html http.FileSystem) http.Handler {
	r := chi.NewRouter()

	hh := handler.NewHelloHander()
	r.Get("/*", http.FileServer(html).ServeHTTP)
	r.Get("/hello/{name}", hh.Hello)

	return r
}
