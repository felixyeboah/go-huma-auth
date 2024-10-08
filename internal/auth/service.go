package auth

import (
	"context"
	"huma-auth/pkg/utils"
)

type Service struct {
	Repository *Repository
}

func NewService(Repository *Repository) *Service {
	return &Service{
		Repository: Repository,
	}
}

func (svc *Service) RegisterUser(ctx context.Context, user UserRequest) (*UserResponse, error) {
	//hash user password
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	// pass the hashedPassword down to the repository
	user.Password = hashedPassword
	u, err := svc.Repository.CreateUser(ctx, user)
	//return user
	return &UserResponse{
		ID:          u.ID,
		Name:        u.Name,
		Email:       u.Email,
		PhoneNumber: u.PhoneNumber,
		IsVerified:  u.IsVerified,
		RoleID:      u.RoleID,
		CreatedAt:   u.CreatedAt,
		UpdatedAt:   u.UpdatedAt,
	}, nil
}

func (svc *Service) VerifyUser(ctx context.Context, userId, token string) error {
	return svc.Repository.VerifyUser(ctx, userId, token)
}

func (svc *Service) LoginUser(ctx context.Context, args LoginUserRequest) (*LoginResponse, error) {
	return svc.Repository.LoginUser(ctx, args)
}

func (svc *Service) ForgotPassword(ctx context.Context, email string) error {
	return svc.Repository.ForgotPassword(ctx, email)
}

func (svc *Service) ResetPassword(ctx context.Context, args ResetPasswordRequest) error {
	return svc.Repository.ResetPassword(ctx, args)
}
