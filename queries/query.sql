-- name: CreateAccount :one
INSERT INTO "Account" (username, "password", balance)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetAccounts :many
SELECT
    a.*,
    json_agg(r.name) AS roles
FROM
    "Account" AS a
JOIN
    "AccountRole" AS ar ON a.username = ar.account
JOIN
    "Role" AS r ON ar.role = r.name
GROUP BY
    a.username
OFFSET $1 LIMIT $2;

-- name: UpdateAccount :exec
UPDATE "Account"
SET username=$1, "password"=$2, balance=$3
WHERE username=$4;

-- name: ReplaceUsername :exec
UPDATE "Account"
SET username=$1
WHERE username=$2;

-- name: AppendRoleAccount :one
INSERT INTO "AccountRole" (account, "role")
VALUES ($1, $2)
RETURNING *;

-- name: DeleteAccountRoles :exec
DELETE
FROM "AccountRole"
WHERE account=$1;


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
