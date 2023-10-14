-- name: CreateAccount :one
INSERT INTO "Account" (username, "password", balance)
VALUES ($1, $2, $3)
RETURNING *;

-- name: AppendRoleAccount :one
INSERT INTO "AccountRole" (account, "role")
VALUES ($1, $2)
RETURNING *;

-- name: IsAccountExist :one
SELECT EXISTS (
  SELECT 1
  FROM "Account"
  WHERE username=$1
);

-- name: GetAccount :one
SELECT *
FROM "Account"
WHERE username=$1;

-- name: GetAccountRoles :many
SELECT "role"
FROM "AccountRole"
WHERE account=$1;

-- name: GetCacheRoles :many
SELECT *
FROM "AccountRole";
