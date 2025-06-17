package httpserver

import (
	"context"
	"net"
	"net/http"

	"github.com/runquan-ray-zhou/uberfx-server/httphandler/api/linknyc"
	"github.com/runquan-ray-zhou/uberfx-server/httphandler/api/pocketdictionary"
	"github.com/runquan-ray-zhou/uberfx-server/httphandler/api/quizme"
	"github.com/runquan-ray-zhou/uberfx-server/httphandler/api/rest"
	"github.com/runquan-ray-zhou/uberfx-server/httphandler/api/rrunquanzhou"
	"github.com/runquan-ray-zhou/uberfx-server/middleware/cors"
	"go.uber.org/fx"

	"go.uber.org/zap"
)

var Module = fx.Options(
	fx.Provide(
		http.NewServeMux,
		newAPIHandlers,
		NewHTTPServer,
		zap.NewProduction,
	),
	fx.Invoke(NewHTTPServer),
)

func NewHTTPServer(lc fx.Lifecycle, mux *http.ServeMux, apiHandlers *APIHandlers, log *zap.Logger) *http.Server {
	RegisterAPIHandlers(mux, apiHandlers)
	handlerWithCORS := cors.CheckCORS(mux)
	srv := &http.Server{Addr: ":8080", Handler: handlerWithCORS}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", srv.Addr)
			if err != nil {
				return err
			}
			log.Info("Starting HTTP server", zap.String("addr", srv.Addr))
			go srv.Serve(ln)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})
	return srv
}

type APIHandlers struct {
	LinkNYCHandler          http.Handler
	PocketDictionaryHandler http.Handler
	QuizMeHandler           http.Handler
	RestHandler             http.Handler
	RayHandler              http.Handler
}

func newAPIHandlers(
	linkNYCHandler *linknyc.Handler,
	pocketDictionaryHandler *pocketdictionary.Handler,
	quizMeHandler *quizme.Handler,
	restHandler *rest.Handler,
	rayHandler *rrunquanzhou.Handler) *APIHandlers {
	return &APIHandlers{
		LinkNYCHandler:          linkNYCHandler,
		PocketDictionaryHandler: pocketDictionaryHandler,
		QuizMeHandler:           quizMeHandler,
		RestHandler:             restHandler,
		RayHandler:              rayHandler,
	}
}

func RegisterAPIHandlers(mux *http.ServeMux, handlers *APIHandlers) {
	mux.Handle("/linknyc", handlers.LinkNYCHandler)
	mux.Handle("/pocketdictionary", handlers.PocketDictionaryHandler)
	mux.Handle("/quizme", handlers.QuizMeHandler)
	mux.Handle("/ray", handlers.RayHandler)
	mux.Handle("/", handlers.RestHandler)
}
