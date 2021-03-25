package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/jsonapi"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// All API Routes
func API(r *mux.Router) *mux.Router {
	log.Println("Loading API routes")

	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {

		// Override the default header middleware
		w.Header().Set("Content-Type", "application/json")

		err := json.NewEncoder(w).Encode(true)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	}).Methods("GET")

	// Route to test CORS headers
	api.Handle("/cors", http.NotFoundHandler()).Methods("GET")

	api.Use(handlers.CORS())
	api.Use(defaultContentType)

	return r

}

func defaultContentType(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", jsonapi.MediaType)
		next.ServeHTTP(w, r)
	})
}
