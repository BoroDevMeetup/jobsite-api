package controllers

import (
	"net/http"
)

func Homepage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Homepage Endpoint Hit🎉"))
}

func NotImplemented(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Not implemented"))
}
