package main

import (
	"net/http"
)

func (apiCfg *apiConfig) GetFeedsHandler(w http.ResponseWriter, r *http.Request) {
	feeds, err := apiCfg.DB.GetFeeds(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't get feeds in DB")
	}
	respondWithJSON(w, http.StatusOK, feeds)
}
