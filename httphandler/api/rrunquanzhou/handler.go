package rrunquanzhou

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
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, `{"message": "Runquan (Ray) Zhou's Website Handler"}`)
}
