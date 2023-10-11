-- name: CreateAccount :one
INSERT INTO "Account" (username, password, "isAdmin", balance)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: IsAccountExist :one
SELECT EXISTS (
  SELECT 1
  FROM "Account"
  WHERE username=$1
);
