package quizme

import (
	"fmt"
	"io"
	"net/http"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

var Module = fx.Provide(
	NewHandler,
)

type Params struct {
	fx.In

	Zap *zap.Logger // logging
	// calling client
	// database
	// auth
}

type Handler struct {
	logger *zap.Logger
}

func NewHandler(p Params) *Handler {
	return &Handler{p.Zap}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if _, err := io.Copy(w, r.Body); err != nil {
		h.logger.Warn("Failed to handle request", zap.Error(err))
	}
	w.Header().Set("Set-Cookie", "somekey=somevalue") //Cookie sample
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, `{"message": "Quiz-Me Handler"}`)
}
