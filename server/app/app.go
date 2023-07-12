package app

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type App struct {
	Router *chi.Mux
}

type optFunc func(app *App)

func (app *App) Start() {
	http.ListenAndServe(":3000", app.Router)
}

func CreateApp(optFuncs ...optFunc) *App {
	app := &App{
		Router: chi.NewRouter(),
	}

	app.Router.Use(middleware.Logger)
	app.Router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Sunset-Cafe backend"))
	})

	for _, optFunc := range optFuncs {
		optFunc(app)
	}

	return app
}
