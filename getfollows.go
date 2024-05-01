package main

import (
	"log"
	"michafdlr/blog_aggregator/internal/database"
	"net/http"
)

func (apiCfg *apiConfig) GetFollowsHandler(w http.ResponseWriter, r *http.Request, user database.User) {
	feeds, err := apiCfg.DB.GetFollows(r.Context(), user.ID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't get feedfollows for user")
	}
	log.Print("returning feedfollows")
	respondWithJSON(w, http.StatusCreated, feeds)
}
