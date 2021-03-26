package routes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/thanhpk/randstr"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/endpoints"
)

var (
	oauthState = randstr.Hex(16)
	authConfig AuthSettings
)

type AuthSettings struct {
	Origin            string
	SlackClientID     string
	SlackClientSecret string
}

func Auth(r *mux.Router) *mux.Router {
	log.Println("Loading auth routes")

	r.HandleFunc("/auth/redirect", func(w http.ResponseWriter, r *http.Request) {

		authConfig := getConfig()
		url := authConfig.AuthCodeURL(oauthState)

		http.Redirect(w, r, url, http.StatusFound)

	})

	r.HandleFunc("/auth/callback", func(w http.ResponseWriter, r *http.Request) {
		code := r.FormValue("code")
		state := r.FormValue("state")

		if state != oauthState {
			http.Error(w, "invalid oauth state", http.StatusInternalServerError)
			return
		}

		authConfig := getConfig()
		token, err := authConfig.Exchange(oauth2.NoContext, code)
		if err != nil {
			fmt.Fprintf(w, "Error: %s", err.Error())
			return
		}

		log.Println(token)

		// teamID := token.Extra("team_id")
		// user := token.Extra("user")
		// fmt.Fprintf(w, "%s %v", teamID, user)

		// err = json.NewEncoder(w).Encode(basicInfo)
		// if err != nil {
		// 	http.Error(w, err.Error(), http.StatusInternalServerError)
		// }

	})

	return r
}

func SetAuthConfig(c AuthSettings) {
	authConfig = c
}

func getConfig() *oauth2.Config {
	origin := authConfig.Origin
	slackClientID := authConfig.SlackClientID
	slackClientSecret := authConfig.SlackClientSecret

	return &oauth2.Config{
		RedirectURL:  origin + "/auth/callback",
		ClientID:     slackClientID,
		ClientSecret: slackClientSecret,
		Scopes:       []string{"identity.basic"},
		Endpoint:     endpoints.Slack,
	}
}
