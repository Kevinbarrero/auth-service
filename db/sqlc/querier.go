// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0

package db

import (
	"context"
)

type Querier interface {
	CreateUser(ctx context.Context, arg CreateUserParams) (Users, error)
	GetUserByEmail(ctx context.Context, email string) (Users, error)
}

var _ Querier = (*Queries)(nil)
