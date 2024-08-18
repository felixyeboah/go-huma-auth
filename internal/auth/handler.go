package auth

import (
	"context"
	"github.com/danielgtaylor/huma/v2"
	"huma-auth/pkg/database"
	"net/http"
)

type RegisterOutput struct {
	Body UserResponse
}

type RegisterInput struct {
	Body UserRequest
}

func RegisterHandlers(api huma.API) {
	//auth
	authRepo := NewRepository(database.Database)
	authService := NewService(authRepo)

	huma.Register(api, huma.Operation{
		OperationID:   "register-user",
		Method:        http.MethodPost,
		Path:          "/api/v1/auth/register",
		Summary:       "Register a user",
		Description:   "Register a user accepting requests",
		Tags:          []string{"Register"},
		DefaultStatus: http.StatusCreated,
	}, func(ctx context.Context, input *RegisterInput) (*RegisterOutput, error) {
		u, err := authService.RegisterUser(ctx, input.Body)
		if err != nil {
			return nil, huma.Error400BadRequest(err.Error(), err)
		}

		createdUser := UserResponse{
			ID:          u.ID,
			Name:        u.Name,
			Email:       u.Email,
			PhoneNumber: u.PhoneNumber,
			IsVerified:  u.IsVerified,
			RoleID:      u.RoleID,
			CreatedAt:   u.CreatedAt,
			UpdatedAt:   u.UpdatedAt,
		}

		return &RegisterOutput{Body: createdUser}, nil
	})
}
