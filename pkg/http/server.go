package server

import (
	"context"
	"net"
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func NewHTTPServer(lc fx.Lifecycle, mux *mux.Router, logger *zap.Logger) *http.Server {
	server := &http.Server{Addr: ":8080", Handler: mux}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", server.Addr)
			if err != nil {
				return err
			}
			logger.Info("Starting a server at ", zap.String("addr", server.Addr))
			go server.Serve(ln)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return server.Shutdown(ctx)
		},
	})
	return server
}
