package auth

import (
	"context"
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"huma-auth/internal/session"
	"huma-auth/pkg/redis"
	"huma-auth/pkg/resend"
	"huma-auth/pkg/token"
	"huma-auth/pkg/utils"
	"huma-auth/sql/sqlc"
	"log"
	"time"
)

type Repository struct {
	Queries    *db.Queries
	db         *sql.DB
	tokenMaker *token.PasetoMaker
	redis      *redis.Store
	session    *session.Repository
}

func NewRepository(database *sql.DB, tokenMaker *token.PasetoMaker, redis *redis.Store,
	session *session.Repository) *Repository {
	return &Repository{
		Queries:    db.New(database),
		db:         database,
		tokenMaker: tokenMaker,
		redis:      redis,
		session:    session,
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
			return errors.New("user with this email already exists")
		} else if existingUser.Email != "" {
			return errors.New("user already exists")
		}

		// Check if role exists
		role, err := queries.GetRoleByName(ctx, args.RoleName)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return errors.New("role not found")
			}
			return errors.New("error getting role")
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
			return errors.New("failed to create verification token")
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

// VerifyUser verify user with their email
func (repo *Repository) VerifyUser(ctx context.Context, userId, token string) error {
	// get the token from redis with the userID
	verificationToken, err := repo.redis.GetToken(ctx, userId)
	if err != nil {
		return errors.New("failed to get token")
	}

	// check if verification token is same as token from the param
	if verificationToken != token {
		return errors.New("invalid token")
	}

	// verify the token with paseto and if it's valid, we move on
	_, err = repo.tokenMaker.VerifyToken(token)
	if err != nil {
		return errors.New("failed to verify token")
	}

	//parse user id from string to uuid
	userID := uuid.MustParse(userId)

	// now, verify user with the id
	err = repo.Queries.VerifyUser(ctx, userID)
	if err != nil {
		return errors.New("failed to verify user")
	}

	// delete the token from redis
	err = repo.redis.DeleteToken(ctx, userId)
	if err != nil {
		return errors.New("failed to delete token")
	}

	return nil
}

func (repo *Repository) LoginUser(ctx context.Context, args LoginUserRequest) (*LoginResponse, error) {
	// fetch user with email
	user, err := repo.Queries.GetUserByEmail(ctx, args.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	err = utils.CheckPasswordHash(user.Password, args.Password)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	// get the user role
	role, err := repo.Queries.GetRole(ctx, user.RoleID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("role not found")
		}

		return nil, err
	}

	// create a session for the user
	s, err := repo.session.CreateSession(ctx, session.Payload{
		UserId:    user.ID,
		UserAgent: args.UserAgent,
		Role:      role.Name,
		IPAddress: args.IPAddress,
	})
	if err != nil {
		return nil, errors.New("failed to create session")
	}

	return &LoginResponse{
		AccessToken:  s.AccessToken,
		RefreshToken: s.RefreshToken,
		User: UserResponse{
			ID:          user.ID,
			Name:        user.Name,
			Email:       user.Email,
			PhoneNumber: user.PhoneNumber,
			RoleID:      user.RoleID,
			IsVerified:  user.IsVerified,
			CreatedAt:   user.CreatedAt,
			UpdatedAt:   user.UpdatedAt,
		},
	}, nil
}

func (repo *Repository) ForgotPassword(ctx context.Context, email string) error {
	// fetch user from db
	user, err := repo.Queries.GetUserByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errors.New("user not found")
		}

		return errors.New("failed to get user")
	}

	// get user role
	role, err := repo.Queries.GetRole(ctx, user.RoleID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errors.New("role not found")
		}

		return errors.New("failed to get role")
	}

	// generate reset token
	t, err := repo.tokenMaker.CreateToken(user.ID.String(), role.Name, 15*time.Minute)
	if err != nil {
		return errors.New("failed to create token")
	}

	// store token in redis
	err = repo.redis.StoreToken(ctx, user.ID.String(), t, 15*time.Minute)
	if err != nil {
		return errors.New("failed to store token")
	}

	// email the user
	err = resend.SendVerificationEmail(user.Email, user.ID.String(), t, "auth/forgot-password")
	if err != nil {
		return errors.New("failed to send verification email")
	}

	return nil
}
