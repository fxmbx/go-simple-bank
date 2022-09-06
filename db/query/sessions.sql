-- name: CreateSession :one
INSERT INTO sessions (
  id,
  username,
  refresh_token,
  user_agent,
  client_ip,
  is_blocked,
  expires_at
) VALUES (
  $1, $2, $3, $4, $5, $6, $7
)
RETURNING *;

-- name: GetSession :one
SELECT * FROM sessions
WHERE id = $1 LIMIT 1;

-- name: GetSessionByUserName :one
SELECT * FROM sessions
WHERE username = $1
LIMIT 1;


-- -- name: ListActiveSessions :many
-- SELECT * FROM sessions
-- WHERE is_blocked = false
-- ORDER BY id 
-- LIMIT $1
-- OFFSET $2;

-- name: UpdateSession :one
UPDATE sessions
set is_blocked = $2
WHERE id = $1
RETURNING *;



-- name: DeleteSession :exec
DELETE FROM sessions
WHERE id = $1;