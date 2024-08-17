package auth

import (
	"context"
	"huma-auth/pkg/utils"
	db "huma-auth/sql/sqlc"
)

type Service struct {
	Repository *Repository
}

func NewService(Repository *Repository) *Service {
	return &Service{
		Repository: Repository,
	}
}

func (svc *Service) RegisterUser(ctx context.Context, user User) (*db.User, error) {
	//hash user password
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	// pass the hashedPassword down to the repository
	user.Password = hashedPassword
	u, err := svc.Repository.CreateUser(ctx, user)
	//return user
	return &db.User{
		ID:          u.ID,
		Name:        u.Name,
		Email:       u.Email,
		PhoneNumber: u.PhoneNumber,
		IsVerified:  u.IsVerified,
	}, nil
}
