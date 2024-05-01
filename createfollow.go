package main

import (
	"encoding/json"
	"log"
	"michafdlr/blog_aggregator/internal/database"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func (apiCfg *apiConfig) CreateFollowHandler(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}

	feedFollow, err := apiCfg.DB.CreateFollow(r.Context(), database.CreateFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.FeedID,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create feed")
		return
	}
	log.Print("New feedfollow created in DB")
	respondWithJSON(w, http.StatusCreated, databaseFollowToFollow(feedFollow))
}
