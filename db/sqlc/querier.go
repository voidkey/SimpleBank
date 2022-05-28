// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0

package db

import (
	"context"
	"database/sql"
)

type Querier interface {
	CreateAccount(ctx context.Context, arg CreateAccountParams) (sql.Result, error)
	GetAccount(ctx context.Context, id int64) (Account, error)
}

var _ Querier = (*Queries)(nil)
