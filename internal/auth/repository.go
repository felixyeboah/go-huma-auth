package auth

import (
	"context"
	"database/sql"
	"errors"
	"huma-auth/pkg/redis"
	"huma-auth/pkg/resend"
	"huma-auth/pkg/token"
	"huma-auth/sql/sqlc"
	"log"
	"time"
)

type Repository struct {
	*db.Queries
	db         *sql.DB
	tokenMaker *token.PasetoMaker
	redis      *redis.Store
}

func NewRepository(database *sql.DB, tokenMaker *token.PasetoMaker, redis *redis.Store) *Repository {
	return &Repository{
		Queries:    db.New(database),
		db:         database,
		tokenMaker: tokenMaker,
		redis:      redis,
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

		// generate verification token
		verificationToken, err := repo.tokenMaker.CreateToken(createdUser.ID.String(), args.RoleName, 15*time.Minute)
		if err != nil {
			return err
		}

		// store verification token in redis
		err = repo.redis.StoreToken(ctx, createdUser.ID.String(), verificationToken, 15*time.Minute)
		if err != nil {
			return errors.New("failed to store verification token")
		}

		// send email with the token
		err = resend.SendVerificationEmail(createdUser.Email, createdUser.ID.String(), verificationToken, "auth/verify")
		if err != nil {
			return errors.New("failed to send verification email")
		}
		return nil
	})

	if err != nil {
		log.Printf("error creating user: %v", err)
		return &UserResponse{}, err
	}

	return createdUser, nil
}
