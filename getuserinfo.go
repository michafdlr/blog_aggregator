package main

import (
	"log"
	"net/http"
)

func (apiCfg *apiConfig) GetCurrentUserHandler(w http.ResponseWriter, r *http.Request) {
	apiKey, err := GetHeader(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "apikey missing")
		return
	}

	user, err := apiCfg.DB.GetUserByKey(r.Context(), apiKey)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't find user")
		return
	}
	log.Print("User info found")
	respondWithJSON(w, http.StatusCreated, databaseUserToUser(user))
}
