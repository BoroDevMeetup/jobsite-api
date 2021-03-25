package routes

import (
	"log"
	"main/controllers"

	"github.com/gorilla/mux"
)

func Home(r *mux.Router) *mux.Router {
	log.Println("Loading home routes")

	r.HandleFunc("/", controllers.Homepage)

	// For now...
	r.HandleFunc("/companies", controllers.CompaniesIndex)
	r.HandleFunc("/people", controllers.PeopleIndex)
	r.HandleFunc("/people/show/{id}", controllers.NotImplemented)

	return r
}
