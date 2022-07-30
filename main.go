package main

import (
	"context"
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net"
	"net/http"

	"github.com/nesheep/go-react-docker-heroku/config"
	"github.com/nesheep/go-react-docker-heroku/server"
)

//go:embed frontend/build
var frontend embed.FS

func main() {
	if err := run(context.Background()); err != nil {
		log.Printf("failed to terminate server: %v", err)
	}
}

func run(ctx context.Context) error {
	html, err := fs.Sub(frontend, "frontend/build")
	if err != nil {
		return err
	}

	cfg, err := config.New()
	if err != nil {
		return err
	}

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		log.Fatalf("failed to listen port %d: %v", cfg.Port, err)
	}

	url := fmt.Sprintf("http://%s", l.Addr().String())
	log.Printf("start with: %v", url)
	log.Printf("env: %v", cfg.Env)

	r := server.NewRouter(cfg, http.FS(html))
	s := server.NewServer(r, l)

	return s.Run(ctx)
}
