package rest

import (
	"fmt"
	"net/http"

	"go.uber.org/fx"
)

var Module = fx.Provide(
	NewHandler,
	NewServeMux,
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (*Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "The Multi App Central Backend Server is Running")
}

func NewServeMux(h *Handler) *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/", h)
	return mux
}
