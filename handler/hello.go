package handler

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Hello struct{}

func NewHello() *Hello {
	return &Hello{}
}

type helloRsp struct {
	Message string `json:"message"`
}

func (hh *Hello) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	name := chi.URLParam(r, "name")
	rsp := helloRsp{Message: fmt.Sprintf("Hello, %v!", name)}
	RespondJSON(ctx, w, rsp, http.StatusOK)
}
