-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, name)
VALUES ($1,$2,$3,$4)
RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE name = $1;

-- name: GetUserById :one
SELECT name FROM users
WHERE id = $1;

-- name: UserExist :one
SELECT EXISTS (
  SELECT 1 from users WHERE name = $1
);

-- name: DeleteUsers :exec
TRUNCATE users RESTART IDENTITY CASCADE;

-- name: GetUsers :many
SELECT * FROM users;
