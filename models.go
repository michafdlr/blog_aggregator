package main

import (
	"michafdlr/blog_aggregator/internal/database"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	ApiKey    string    `json:"api_key"`
}

type Feed struct {
	ID            uuid.UUID
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Name          string
	Url           string
	UserID        uuid.UUID
	LastFetchedAt time.Time
}

type FeedsFollow struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    uuid.UUID
	FeedID    uuid.UUID
}

func databaseUserToUser(user database.User) User {
	return User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Name:      user.Name,
		ApiKey:    user.ApiKey,
	}
}

// func databaseFeedToFeed(feed Feed) Feed {
// 	return Feed{
// 		ID:            feed.ID,
// 		CreatedAt:     feed.CreatedAt,
// 		UpdatedAt:     feed.UpdatedAt,
// 		Name:          feed.Name,
// 		Url:           feed.Url,
// 		UserID:        feed.UserID,
// 		LastFetchedAt: feed.LastFetchedAt,
// 	}
// }

func databaseFollowToFollow(feedFollow database.FeedsFollow) FeedsFollow {
	return FeedsFollow{
		ID:        feedFollow.ID,
		CreatedAt: feedFollow.CreatedAt,
		UpdatedAt: feedFollow.UpdatedAt,
		UserID:    feedFollow.UserID,
		FeedID:    feedFollow.FeedID,
	}
}
