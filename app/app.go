package app

import (
	"log"
	"main/routes"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

type App struct {
	Settings *viper.Viper
	Mux      *mux.Router
}

func (a *App) Run() {
	log.Println("Running app")

	addr := a.Settings.Get("addr").(string)

	a.LoadRoutes()
	lr := handlers.LoggingHandler(os.Stdout, a.Mux)

	log.Print("Setup server on: ", addr)
	log.Fatal(http.ListenAndServe(addr, lr))
}

func (a *App) Close() {
	log.Println("Closing down app")
}

func (a *App) LoadRoutes() {
	origin := a.Settings.Get("origin").(string)
	slackClientID := a.Settings.Get("slack_client_id").(string)
	slackClientSecret := a.Settings.Get("slack_client_secret").(string)

	a.Mux = routes.Home(a.Mux)
	a.Mux = routes.API(a.Mux)

	s := routes.AuthSettings{
		Origin:            origin,
		SlackClientID:     slackClientID,
		SlackClientSecret: slackClientSecret,
	}

	routes.SetAuthConfig(s)
	a.Mux = routes.Auth(a.Mux)
}
