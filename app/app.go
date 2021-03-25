package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

type App struct {
	Settings  *viper.Viper
	Mux       *mux.Router
	Templates *template.Template
}

func (a *App) Run() {
	log.Println("Running app")

	addr := viper.Get("addr").(string)

	// r := mux.NewRouter().StrictSlash(true)

	// r = routes.Home(r)
	// r = routes.API(r)
	// routes.SetAppSettings(&a)
	// r = routes.Auth(r)
	// a.SetRoutes(r)
	// r = routes.API(r)

	// routes.SetAppSettings(a)
	// r = routes.Auth(r)

	// a.Mux = r

	lr := handlers.LoggingHandler(os.Stdout, a.Mux)
	log.Print("Setup server on: ", addr)
	log.Fatal(http.ListenAndServe(addr, lr))
}

func (a *App) Close() {
	log.Println("Closing down app")
}

func (a *App) SetRoutes(r *mux.Router) {
	a.Mux = r
}

func (a *App) LoadSettings() {

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	a.Settings = viper.GetViper()
}
