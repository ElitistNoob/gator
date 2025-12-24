-- name: CreatePost :one
INSERT INTO posts(id, created_at, updated_at, title, url, description, published_at, feed_id)
VALUES($1,$2,$3,$4,$5,$6,$7,$8)
RETURNING *;

-- name: GetPostsForUser :many
WITH users_feed AS (
  SELECT feed_id FROM feed_follows
  WHERE feed_follows.user_id = $1
)

SELECT 
  posts.id,
  posts.created_at, 
  posts.updated_at, 
  posts.title, 
  posts.url, 
  posts.description, 
  posts.published_at, 
  posts.feed_id, 
  f.name AS feed_name 
FROM posts
JOIN users_feed uf ON posts.feed_id = uf.feed_id
JOIN feeds f ON posts.feed_id = f.id
WHERE posts.published_at >= $2 AND posts.published_at < $3
ORDER BY 
CASE WHEN $4 = 'asc' THEN posts.published_at END ASC,
CASE WHEN $4 = 'desc' THEN posts.published_at END DESC
LIMIT $5;
