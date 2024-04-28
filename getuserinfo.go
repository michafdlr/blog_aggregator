package main

import (
	"michafdlr/blog_aggregator/internal/database"
	"net/http"
)

func (apiCfg *apiConfig) GetCurrentUserHandler(w http.ResponseWriter, r *http.Request, user database.User) {
	respondWithJSON(w, http.StatusCreated, databaseUserToUser(user))
}
