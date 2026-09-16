-- name: CreateURL :one
INSERT INTO urls (
    short_code,
    destination
) VALUES ($1, $2) RETURNING *;

-- name: GetURLFromShortCode :one
SELECT
    *
FROM
    urls
WHERE
    short_code = $1;
