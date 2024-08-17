package auth

import (
	"context"
	"database/sql"
	"errors"
	db "huma-auth/sql/sqlc"
)

type Repository struct {
	Queries db.Querier
}

func NewRepository(queries db.Querier) Repository {
	return Repository{
		Queries: queries,
	}
}

// CreateUser adds a new user
func (repo *Repository) CreateUser(ctx context.Context, user User) (*db.User, error) {
	// Check if user exists
	existingUser, err := repo.Queries.GetUserByEmail(ctx, user.Email)
	if err == nil && errors.Is(err, sql.ErrNoRows) {
		return nil, err
	} else if existingUser.Email != "" {
		return nil, errors.New("user already exists")
	}

	// Check if role exists
	role, err := repo.Queries.GetRoleByName(ctx, user.RoleName)
	if err != nil {
		return nil, err
	}
	// Create a user
	u, err := repo.Queries.CreateUser(ctx, db.CreateUserParams{
		Name:        user.Name,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		Password:    user.Password,
		RoleID:      role.ID,
	})
	if err != nil {
		return nil, errors.New("failed to create user")
	}

	// TODO: generate a validation token, save in redis and send to the user via email

	return &db.User{
		ID:          u.ID,
		Name:        u.Name,
		Email:       u.Email,
		PhoneNumber: u.PhoneNumber,
		RoleID:      role.ID,
		IsVerified:  u.IsVerified,
	}, nil
}
