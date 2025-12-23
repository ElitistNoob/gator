-- name: CreatePost :one
INSERT INTO posts(id, created_at, updated_at, title, url, description, published_at, feed_id)
VALUES($1,$2,$3,$4,$5,$6,$7,$8)
RETURNING *;

-- name: GetPostsForUser :many
WITH users_feed AS (
  SELECT feed_id FROM feed_follows
  WHERE feed_follows.user_id = $1
)

SELECT posts.*, f.name AS feed_name FROM posts
JOIN users_feed uf ON posts.feed_id = uf.feed_id
JOIN feeds f ON posts.feed_id = f.id
ORDER BY posts.created_at DESC
LIMIT $2;
