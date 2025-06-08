package rest

import (
	"fmt"
	"io"
	"net/http"
	"os"

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
	if _, err := io.Copy(w, r.Body); err != nil {
		fmt.Fprintln(os.Stderr, "Failed to handle request:", err)
	}
}

func NewServeMux(h *Handler) *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/", h)
	return mux
}
