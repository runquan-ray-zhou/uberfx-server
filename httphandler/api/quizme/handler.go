package quizme

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
	w.Header().Set("Set-Cookie", "somekey=somevalue") //Cookie sample
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, `{"message": "Quiz-Me Handler"}`)
}
