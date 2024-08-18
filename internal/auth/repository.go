package auth

import (
	"context"
	"database/sql"
	"errors"
	"huma-auth/sql/sqlc"
	"log"
)

type Repository struct {
	*db.Queries
	db *sql.DB
}

func NewRepository(database *sql.DB) *Repository {
	return &Repository{
		Queries: db.New(database),
		db:      database,
	}
}

// CreateUser adds a new user
func (repo *Repository) CreateUser(ctx context.Context, args UserRequest) (*UserResponse, error) {
	var createdUser *UserResponse

	// Check if user exists
	err := db.ExecTX(ctx, repo.db, func(tx *sql.Tx) error {
		if tx == nil {
			return errors.New("transaction is nil")
		}
		queries := db.New(tx) // Create a new queries instance with the transaction

		existingUser, err := queries.GetUserByEmail(ctx, args.Email)
		if err == nil && errors.Is(err, sql.ErrNoRows) {
			return err
		} else if existingUser.Email != "" {
			return errors.New("user already exists")
		}

		// Check if role exists
		role, err := queries.GetRoleByName(ctx, args.RoleName)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return errors.New("role not found")
			}
			return err
		}

		// Create a user
		createdUserDB, err := queries.CreateUser(ctx, db.CreateUserParams{
			Name:        args.Name,
			Email:       args.Email,
			PhoneNumber: args.PhoneNumber,
			Password:    args.Password,
			RoleID:      role.ID,
		})
		if err != nil {
			return errors.New("failed to create user")
		}

		createdUser = &UserResponse{
			ID:          createdUserDB.ID,
			Name:        createdUserDB.Name,
			Email:       createdUserDB.Email,
			PhoneNumber: createdUserDB.PhoneNumber,
			RoleID:      createdUserDB.RoleID,
			IsVerified:  createdUserDB.IsVerified,
			CreatedAt:   createdUserDB.CreatedAt,
			UpdatedAt:   createdUserDB.UpdatedAt,
		}

		// TODO: generate a validation token, save in redis and send to the user via email

		return nil
	})

	if err != nil {
		log.Printf("error creating user: %v", err)
		return &UserResponse{}, err
	}

	return createdUser, nil
}
