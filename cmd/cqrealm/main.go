package main

import (
	handlers "mucahiderenler/conquerors-realm/internal/handlers"
	repos "mucahiderenler/conquerors-realm/internal/repository"
	services "mucahiderenler/conquerors-realm/internal/services"
	db "mucahiderenler/conquerors-realm/pkg/db"
	server "mucahiderenler/conquerors-realm/pkg/http"
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func main() {
	fx.New(
		fx.Provide(server.NewHTTPServer,
			fx.Annotate(
				NewServeMux,
				fx.ParamTags(`group:"routes"`),
			)),
		fx.Provide(
			AsRoute(handlers.NewVillageHandler),
			services.NewVillageService,
			repos.NewVillageRepository,
			AsRoute(handlers.NewMapHandler),
			services.NewMapService,
			repos.NewMapRepository,
			db.ProvideDB,
			zap.NewExample,
		),
		fx.Invoke(func(*http.Server) {}),
	).Run()
}

func NewServeMux(routes []Route) *mux.Router {
	mux := mux.NewRouter()
	for _, route := range routes {
		route.RegisterRoutes(mux)
	}
	return mux
}

type Route interface {
	RegisterRoutes(*mux.Router)
}

func AsRoute(f any) any {
	return fx.Annotate(
		f,
		fx.As(new(Route)),
		fx.ResultTags(`group:"routes"`),
	)
}
