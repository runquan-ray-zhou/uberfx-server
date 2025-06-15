package quizme

import (
	"io"
	"net/http"
	"time"

	"github.com/runquan-ray-zhou/uberfx-server/httphandler/client"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

const (
	_openTriviaDataBaseAPIBaseURL = "https://opentdb.com/"
)

var Module = fx.Provide(
	NewHandler,
	NewOpenTDBClient,
)

type Params struct {
	fx.In

	Zap    *zap.Logger       // logging
	Client *client.APIClient // calling client
	// database
	// auth
}

type Handler struct {
	logger    *zap.Logger
	apiClient *client.APIClient
}

func NewHandler(p Params) *Handler {
	return &Handler{
		p.Zap,
		p.Client,
	}
}

// NewOpenTDBClient returns APIClient struct
func NewOpenTDBClient() *client.APIClient {
	return client.NewAPIClient(_openTriviaDataBaseAPIBaseURL, 10*time.Second)
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// response is a random question
	resp, err := h.apiClient.HttpClient.Get(h.apiClient.BaseURL + `api.php?amount=1&category=&difficulty=&type=`)
	if err != nil {
		h.logger.Error("API call failed", zap.Error(err))
		http.Error(w, "API call failed", http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		h.logger.Error("Failed to read API response", zap.Error(err))
		http.Error(w, "Failed to read response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Set-Cookie", "somekey=somevalue") // random cookie
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}
