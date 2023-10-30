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
    "AccountRole" AS ar ON a.id = ar.account
JOIN
    "Role" AS r ON ar.role = r.name
GROUP BY
    a.id
OFFSET $1 LIMIT $2;

-- name: GetExistAccounts :many
SELECT
    a.*,
    json_agg(r.name) AS roles
FROM
    "Account" AS a
JOIN
    "AccountRole" AS ar ON a.id = ar.account
JOIN
    "Role" AS r ON ar.role = r.name
WHERE
    a.deleted = false
GROUP BY
    a.id
OFFSET $1 LIMIT $2;

-- name: UpdateAccount :exec
UPDATE "Account"
SET username=$1, "password"=$2, balance=$3
WHERE id=$4;

-- name: ReplaceUsername :exec
UPDATE "Account"
SET username=$1
WHERE id=$2;

-- name: AppendRoleAccount :one
INSERT INTO "AccountRole" (account, "role")
VALUES ($1, $2)
RETURNING *;

-- name: DeleteAccountRoles :exec
DELETE
FROM "AccountRole"
WHERE account=$1;

-- name: IsAccountExistByID :one
SELECT EXISTS (
  SELECT 1
  FROM "Account"
  WHERE id=$1
);

-- name: IsAccountExistByUsername :one
SELECT EXISTS (
  SELECT 1
  FROM "Account"
  WHERE username=$1
);

-- name: IsAccountRemovedByID :one
SELECT EXISTS (
  SELECT 1
  FROM "Account"
  WHERE id=$1 and deleted=true
);

-- name: IsAccountRemovedByUsername :one
SELECT EXISTS (
  SELECT 1
  FROM "Account"
  WHERE username=$1 and deleted=true
);

-- name: RemoveAccount :exec
UPDATE "Account"
SET deleted=true
WHERE id=$1;

-- name: GetAccountByID :one
SELECT *
FROM "Account"
WHERE id=$1;

-- name: GetAccountByUsername :one
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
("owner", "type", can_rented, model, color, "description", identifier, latitude, longitude, minute_price, day_price, deleted)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, false)
RETURNING *;

-- name: IsExistTransportByID :one
SELECT EXISTS (
  SELECT 1
  FROM "Transport"
  WHERE id=$1
);

-- name: IsExistTransportByIdentifier :one
SELECT EXISTS (
  SELECT 1
  FROM "Transport"
  WHERE identifier=$1
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
SET can_rented=$1, model=$2, color=$3, "description"=$4, latitude=$5, longitude=$6, minute_price=$7, day_price=$8, identifier=$9
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

-- name: GetTransports :many
SELECT *
FROM "Transport"
ORDER BY id
OFFSET $1 LIMIT $2;

-- name: GetTransportsOnlyType :many
SELECT *
FROM "Transport"
WHERE "type"=$1
ORDER BY id
OFFSET $2 LIMIT $3;

-- name: GetTransportsFromRadiusAll :many
SELECT *
FROM "Transport"
WHERE
    "can_rented"=true
    and point(longitude, latitude) <@> point($1, $2) <= $3;

-- name: GetTransportsFromRadiusOnlyType :many
SELECT *
FROM "Transport"
WHERE
    "can_rented"=true
    and "type"=$1
    and point(longitude, latitude) <@> point($2, $3) <= $4;

-- name: IsRentRemoved :one
SELECT EXISTS (
  SELECT 1
  FROM "Rent"
  WHERE id=$1 and deleted=true
);

-- name: IsRentExist :one
SELECT EXISTS (
  SELECT 1
  FROM "Rent"
  WHERE id=$1
);

-- name: IsRenter :one
SELECT EXISTS (
  SELECT 1
  FROM "Rent"
  WHERE id=$1 and account=$2
);

-- name: GetRent :one
SELECT *
FROM "Rent"
WHERE "Rent".id=$1;

-- name: GetRentsFromID :many
SELECT *
FROM "Rent"
WHERE account=$1
OFFSET $2 LIMIT $3;

-- name: GetRentsFromTransportID :many
SELECT *
FROM "Rent"
WHERE transport=$1
OFFSET $2 LIMIT $3;

-- name: IsExistCurrentRent :one
SELECT EXISTS (
    SELECT 1
    FROM "Rent"
    WHERE transport=$1 and time_end=null
);

-- name: CreateRent :one
INSERT INTO "Rent" (account, transport, time_start, time_end, price_unit, price_type, deleted)
VALUES ($1, $2, $3, $4, $5, $6, false)
RETURNING *;

-- name: EndRent :exec
UPDATE "Rent"
SET time_end=$1
WHERE id=$2;

-- name: UpdateRent :one
UPDATE "Rent"
SET account=$1, transport=$2, time_start=$3, time_end=$4, price_unit=$5, price_type=$6
WHERE id=$7
RETURNING *;

-- name: RemoveRent :exec
UPDATE "Rent"
SET deleted=true
WHERE id=$1;
