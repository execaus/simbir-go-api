-- name: CreateAccount :one
INSERT INTO "Account" (username, "password", balance)
VALUES ($1, $2, $3)
RETURNING *;

-- name: UpdateAccount :exec
UPDATE "Account"
SET username=$1, "password"=$2
WHERE username=$3;

-- name: ReplaceUsername :exec
UPDATE "Account"
SET username=$1
WHERE username=$2;

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

-- name: AppendTokenToBlackList :exec
INSERT INTO "TokenBlackList" (token)
VALUES ($1);

-- name: IsContainBlackListToken :one
SELECT EXISTS (
  SELECT 1
  FROM "TokenBlackList"
  WHERE token=$1
);
