-- name: CreateURL :one
INSERT INTO urls (
    short_code,
    destination
) VALUES ($1, $2) RETURNING *;
