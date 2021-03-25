package controllers

import (
	"html/template"
	"net/http"
)

var TPL *template.Template

func init() {
	TPL = template.Must(template.ParseGlob("views/*/*.html"))
}

func Homepage(w http.ResponseWriter, r *http.Request) {
	err := TPL.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func NotImplemented(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Not implemented"))
}
