// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	CreateSession(ctx context.Context, arg CreateSessionParams) (Session, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	GetRole(ctx context.Context, id uuid.UUID) (Role, error)
	GetRoleByName(ctx context.Context, name string) (GetRoleByNameRow, error)
	GetUserByEmail(ctx context.Context, email string) (User, error)
	VerifyUser(ctx context.Context, id uuid.UUID) error
}

var _ Querier = (*Queries)(nil)
