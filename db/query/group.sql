-- name: CreateGroup :one
INSERT INTO groups (
  owner,
  group_name
) VALUES (
  $1, $2
) RETURNING *;

-- name: GetGroup :one
SELECT * FROM groups
WHERE id = $1 LIMIT 1;

-- name: GetGroupForUpdate :one
SELECT * FROM groups
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: ListGroups :many
SELECT * FROM groups
WHERE owner = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: UpdateGroup :one
UPDATE groups
SET group_name = $2
WHERE id = $1
RETURNING *;

-- name: DeactivateGroup :exec
UPDATE groups
SET status = $2
WHERE id = $1
RETURNING *;
