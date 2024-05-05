-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id, last_fetched_at)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: GetFeeds :many
SELECT * FROM feeds;

-- name: CreateFollow :one
INSERT INTO feeds_follows (id, created_at, updated_at, user_id, feed_id)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: DeleteFollow :exec
DELETE FROM feeds_follows
WHERE id = $1;

-- name: GetFollows :many
SELECT * FROM feeds_follows
WHERE user_id = $1;
