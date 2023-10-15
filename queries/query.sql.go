// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: query.sql

package queries

import (
	"context"
	"encoding/json"
)

const appendRoleAccount = `-- name: AppendRoleAccount :one
INSERT INTO "AccountRole" (account, "role")
VALUES ($1, $2)
RETURNING account, role
`

type AppendRoleAccountParams struct {
	Account string
	Role    string
}

func (q *Queries) AppendRoleAccount(ctx context.Context, arg AppendRoleAccountParams) (AccountRole, error) {
	row := q.db.QueryRowContext(ctx, appendRoleAccount, arg.Account, arg.Role)
	var i AccountRole
	err := row.Scan(&i.Account, &i.Role)
	return i, err
}

const appendTokenToBlackList = `-- name: AppendTokenToBlackList :exec
INSERT INTO "TokenBlackList" (token)
VALUES ($1)
`

func (q *Queries) AppendTokenToBlackList(ctx context.Context, token string) error {
	_, err := q.db.ExecContext(ctx, appendTokenToBlackList, token)
	return err
}

const createAccount = `-- name: CreateAccount :one
INSERT INTO "Account" (username, "password", balance)
VALUES ($1, $2, $3)
RETURNING username, password, balance
`

type CreateAccountParams struct {
	Username string
	Password string
	Balance  float64
}

func (q *Queries) CreateAccount(ctx context.Context, arg CreateAccountParams) (Account, error) {
	row := q.db.QueryRowContext(ctx, createAccount, arg.Username, arg.Password, arg.Balance)
	var i Account
	err := row.Scan(&i.Username, &i.Password, &i.Balance)
	return i, err
}

const deleteAccountRoles = `-- name: DeleteAccountRoles :exec
DELETE
FROM "AccountRole"
WHERE account=$1
`

func (q *Queries) DeleteAccountRoles(ctx context.Context, account string) error {
	_, err := q.db.ExecContext(ctx, deleteAccountRoles, account)
	return err
}

const getAccount = `-- name: GetAccount :one
SELECT username, password, balance
FROM "Account"
WHERE username=$1
`

func (q *Queries) GetAccount(ctx context.Context, username string) (Account, error) {
	row := q.db.QueryRowContext(ctx, getAccount, username)
	var i Account
	err := row.Scan(&i.Username, &i.Password, &i.Balance)
	return i, err
}

const getAccountRoles = `-- name: GetAccountRoles :many
SELECT "role"
FROM "AccountRole"
WHERE account=$1
`

func (q *Queries) GetAccountRoles(ctx context.Context, account string) ([]string, error) {
	rows, err := q.db.QueryContext(ctx, getAccountRoles, account)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var role string
		if err := rows.Scan(&role); err != nil {
			return nil, err
		}
		items = append(items, role)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAccounts = `-- name: GetAccounts :many
SELECT
    a.username, a.password, a.balance,
    json_agg(r.name) AS roles
FROM
    "Account" AS a
JOIN
    "AccountRole" AS ar ON a.username = ar.account
JOIN
    "Role" AS r ON ar.role = r.name
GROUP BY
    a.username
OFFSET $1 LIMIT $2
`

type GetAccountsParams struct {
	Offset int32
	Limit  int32
}

type GetAccountsRow struct {
	Username string
	Password string
	Balance  float64
	Roles    json.RawMessage
}

func (q *Queries) GetAccounts(ctx context.Context, arg GetAccountsParams) ([]GetAccountsRow, error) {
	rows, err := q.db.QueryContext(ctx, getAccounts, arg.Offset, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetAccountsRow
	for rows.Next() {
		var i GetAccountsRow
		if err := rows.Scan(
			&i.Username,
			&i.Password,
			&i.Balance,
			&i.Roles,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getCacheRoles = `-- name: GetCacheRoles :many
SELECT account, role
FROM "AccountRole"
`

func (q *Queries) GetCacheRoles(ctx context.Context) ([]AccountRole, error) {
	rows, err := q.db.QueryContext(ctx, getCacheRoles)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []AccountRole
	for rows.Next() {
		var i AccountRole
		if err := rows.Scan(&i.Account, &i.Role); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const isAccountExist = `-- name: IsAccountExist :one
SELECT EXISTS (
  SELECT 1
  FROM "Account"
  WHERE username=$1
)
`

func (q *Queries) IsAccountExist(ctx context.Context, username string) (bool, error) {
	row := q.db.QueryRowContext(ctx, isAccountExist, username)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const isContainBlackListToken = `-- name: IsContainBlackListToken :one
SELECT EXISTS (
  SELECT 1
  FROM "TokenBlackList"
  WHERE token=$1
)
`

func (q *Queries) IsContainBlackListToken(ctx context.Context, token string) (bool, error) {
	row := q.db.QueryRowContext(ctx, isContainBlackListToken, token)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const replaceUsername = `-- name: ReplaceUsername :exec
UPDATE "Account"
SET username=$1
WHERE username=$2
`

type ReplaceUsernameParams struct {
	Username   string
	Username_2 string
}

func (q *Queries) ReplaceUsername(ctx context.Context, arg ReplaceUsernameParams) error {
	_, err := q.db.ExecContext(ctx, replaceUsername, arg.Username, arg.Username_2)
	return err
}

const updateAccount = `-- name: UpdateAccount :exec
UPDATE "Account"
SET username=$1, "password"=$2, balance=$3
WHERE username=$4
`

type UpdateAccountParams struct {
	Username   string
	Password   string
	Balance    float64
	Username_2 string
}

func (q *Queries) UpdateAccount(ctx context.Context, arg UpdateAccountParams) error {
	_, err := q.db.ExecContext(ctx, updateAccount,
		arg.Username,
		arg.Password,
		arg.Balance,
		arg.Username_2,
	)
	return err
}
