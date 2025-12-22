-- name: CreatePost :one
INSERT INTO posts(id, created_at, updated_at, title, url, description, published_at, feed_id)
VALUES($1,$2,$3,$4,$5,$6,$7,$8)
RETURNING *;

-- name: GetPostsForUser :many
WITH users_feed AS (
  SELECT * FROM feed_follows
  WHERE feed_follows.user_id = $1
)

SELECT * FROM posts
INNER JOIN users_feed ON posts.feed_id = users_feed.user_id
ORDER BY posts.created_at DESC
LIMIT $2;
