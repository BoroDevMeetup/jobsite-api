package controllers

import (
	"encoding/json"
	"net/http"
)

func CompaniesIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(true)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
