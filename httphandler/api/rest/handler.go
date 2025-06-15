package rest

import (
	"fmt"
	"net/http"

	"go.uber.org/fx"
)

var Module = fx.Provide(
	NewHandler,
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (*Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, `{"message": "The Multi App Central Backend Server is Running"}`)
}
