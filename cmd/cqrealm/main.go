package main

import (
	handlers "mucahiderenler/conquerors-realm/internal/handlers"
	repos "mucahiderenler/conquerors-realm/internal/repository"
	services "mucahiderenler/conquerors-realm/internal/services"
	db "mucahiderenler/conquerors-realm/pkg/db"
	server "mucahiderenler/conquerors-realm/pkg/http"
	"net/http"

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
			services.NewVillageService,
			repos.NewVillageRepository,
			db.ProvideDB,
			zap.NewExample,
		),
		fx.Invoke(func(*http.Server) {}),
	).Run()
}
