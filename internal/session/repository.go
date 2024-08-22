package session

import (
	"context"
	"database/sql"
	"huma-auth/pkg/token"
	"huma-auth/pkg/utils"
	db "huma-auth/sql/sqlc"
	"time"
)

type Repository struct {
	Queries *db.Queries
	token   *token.PasetoMaker
}

func NewRepository(database *sql.DB, token *token.PasetoMaker) *Repository {
	return &Repository{
		Queries: db.New(database),
		token:   token,
	}
}

func (repo *Repository) CreateSession(ctx context.Context, args Payload) (Response, error) {
	// Set the session expiry date (30days)
	sessionExpiry := time.Now().Add(30 * 24 * time.Hour)

	// generate access token
	accessToken, err := repo.token.CreateToken(args.UserId.String(), args.Role, 24*time.Hour)
	if err != nil {
		return Response{}, err
	}

	refreshToken, err := repo.token.CreateToken(args.UserId.String(), args.Role, 7*24*time.Hour)
	if err != nil {
		return Response{}, err
	}

	session, err := repo.Queries.CreateSession(ctx, db.CreateSessionParams{
		AccessToken:    accessToken,
		RefreshToken:   refreshToken,
		UserID:         args.UserId,
		UserAgent:      utils.ConvertStringToSQLNullString(args.UserAgent),
		IpAddress:      utils.ConvertStringToSQLNullString(args.IPAddress),
		ExpiryDate:     sessionExpiry,
		LastAccessedAt: utils.ConvertTotimeToSQLNullTime(time.Now()),
	})
	if err != nil {
		return Response{}, err
	}

	return Response{
		ID:             session.ID,
		AccessToken:    session.AccessToken,
		RefreshToken:   session.RefreshToken,
		UserID:         session.UserID,
		ExpiryDate:     session.ExpiryDate,
		UserAgent:      utils.ConvertSQLNullStringToString(session.UserAgent),
		IpAddress:      utils.ConvertSQLNullStringToString(session.IpAddress),
		LastAccessedAt: utils.ConvertSQLNullTimeToTime(session.LastAccessedAt),
	}, nil
}
