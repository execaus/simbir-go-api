-- name: CreateAccount :one
INSERT INTO "Account" (username, "password", balance, deleted)
VALUES ($1, $2, $3, false)
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

-- name: GetExistAccounts :many
SELECT
    a.*,
    json_agg(r.name) AS roles
FROM
    "Account" AS a
JOIN
    "AccountRole" AS ar ON a.username = ar.account
JOIN
    "Role" AS r ON ar.role = r.name
WHERE
    a.deleted = false
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

-- name: IsAccountRemoved :one
SELECT EXISTS (
  SELECT 1
  FROM "Account"
  WHERE username=$1 and deleted=true
);

-- name: RemoveAccount :exec
UPDATE "Account"
SET deleted=true
WHERE username=$1;

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

-- name: CreateTransport :one
INSERT INTO "Transport"
(id, "owner", "type", can_ranted, model, color, "description", latitude, longitude, minute_price, day_price, deleted)
VALUES
($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, false)
RETURNING *;

-- name: IsExistTransport :one
SELECT EXISTS (
  SELECT 1
  FROM "Transport"
  WHERE id=$1
);

-- name: GetTransport :one
SELECT *
FROM "Transport"
WHERE id=$1;

-- name: IsTransportOwner :one
SELECT EXISTS (
  SELECT 1
  FROM "Transport"
  WHERE id=$1 and "owner"=$2
);

-- name: UpdateTransport :one
UPDATE "Transport"
SET id=$1, can_ranted=$2, model=$3, color=$4, "description"=$5, latitude=$6, longitude=$7, minute_price=$8, day_price=$9
WHERE id=$10
RETURNING *;

-- name: RemoveTransport :exec
UPDATE "Transport"
SET deleted=true
WHERE id=$1;

-- name: IsTransportRemoved :one
SELECT EXISTS (
  SELECT 1
  FROM "Transport"
  WHERE id=$1 and deleted=true
);
