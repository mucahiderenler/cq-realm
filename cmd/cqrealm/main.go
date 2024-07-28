package main

import (
	handlers "mucahiderenler/conquerors-realm/internal/handlers"
	server "mucahiderenler/conquerors-realm/pkg/http"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

func main() {
	fx.New(
		fx.Provide(server.NewHTTPServer,
			fx.Annotate(
				server.NewServeMux,
				fx.ParamTags(`group:"routes"`),
			)),
		fx.Provide(

			server.AsRoute(handlers.NewVillageHandler),
			zap.NewExample(),
		),
	).Run()
}
