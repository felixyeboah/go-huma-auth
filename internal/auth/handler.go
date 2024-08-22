package auth

import (
	"context"
	"github.com/danielgtaylor/huma/v2"
	"huma-auth/pkg/database"
	"huma-auth/pkg/redis"
	"huma-auth/pkg/token"
	"huma-auth/pkg/utils"
	"net/http"
)

func RegisterHandlers(api huma.API) {
	// init token make
	tokenMaker, err := token.NewPasetoMaker(utils.RandomString(32))
	if err != nil {
		panic(err)
	}

	//init redis
	client := redis.NewRedisClient()
	redisClient := redis.NewStore(client)

	//auth
	authRepo := NewRepository(database.Database, tokenMaker, redisClient)
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
