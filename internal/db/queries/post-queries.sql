-- name: CreatePost :one
INSERT INTO posts (user_id, content) VALUES ($1, $2)
RETURNING id, user_id, content, created_at;

-- name: ListPostsByUser :many
SELECT id, user_id, content, created_at
FROM posts
WHERE user_id = $1
ORDER BY created_at DESC;

-- name: ListPosts :many
SELECT id, user_id, content, created_at
FROM posts
ORDER BY created_at DESC;