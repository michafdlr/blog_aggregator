package main

import (
	"log"
	"net/http"

	"github.com/google/uuid"
)

func (apiCfg *apiConfig) DeleteFollowHandler(w http.ResponseWriter, r *http.Request) {
	feedID, err := uuid.Parse(r.PathValue("feedFollowID"))
	log.Print(feedID)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't parse feedfollowid")
		return
	}

	err = apiCfg.DB.DeleteFollow(r.Context(), feedID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't delete feed")
		return
	}
	log.Print("Deleted feedfollow")
	// respondWithJSON(w, http.StatusCreated, databaseFollowToFollow(feedFollow))
}
