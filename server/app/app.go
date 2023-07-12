package app

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jakubson7/sunset-cafe/db"
)

type App struct {
	Router     *chi.Mux
	DBProvider db.DBProvider
}

type Plugin func(app *App)

func (app *App) Start() {
	http.ListenAndServe(":3000", app.Router)
}

func CreateApp(plugins ...Plugin) *App {
	app := &App{
		Router:     chi.NewRouter(),
		DBProvider: db.CreateSqliteProvider(),
	}

	app.Router.Use(middleware.Logger)
	app.Router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Sunset-Cafe backend"))
	})

	for _, plugin := range plugins {
		plugin(app)
	}

	return app
}
