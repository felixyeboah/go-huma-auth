package auth

import (
	"context"
	"database/sql"
	"github.com/danielgtaylor/huma/v2"
	"huma-auth/internal/session"
	"huma-auth/pkg/redis"
	"huma-auth/pkg/token"
	"huma-auth/pkg/utils"
	"net/http"
)

func (m *LoginUserInput) Resolve(ctx huma.Context) []error {
	m.Body.UserAgent = ctx.Header("User-Agent")
	m.Body.IPAddress = ctx.RemoteAddr()

	return nil
}

func RegisterHandlers(api huma.API, database *sql.DB) {

	// init token make
	tokenMaker, err := token.NewPasetoMaker(utils.RandomString(32))
	if err != nil {
		panic(err)
	}

	//init redis
	client := redis.NewRedisClient()
	redisClient := redis.NewStore(client)

	// init session
	s := session.NewRepository(database, tokenMaker)

	//auth
	authRepo := NewRepository(database, tokenMaker, redisClient, s)
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

		resp := &RegisterOutput{}
		resp.Body.User = createdUser
		resp.Body.Status = http.StatusCreated
		resp.Body.Message = "User successfully created!"

		return resp, nil
	})

	huma.Register(api, huma.Operation{
		OperationID:   "verify-user",
		Method:        http.MethodGet,
		Path:          "/api/v1/auth/verify",
		Summary:       "Verifies a user",
		Description:   "Verifies a user grabbing their user ID and token from the params",
		Tags:          []string{"Verify"},
		DefaultStatus: http.StatusCreated,
	}, func(ctx context.Context, input *VerifyInput) (*VerifyOutput, error) {
		err := authService.VerifyUser(ctx, input.UserId, input.Token)
		if err != nil {
			return nil, huma.Error400BadRequest(err.Error(), err)
		}
		resp := &VerifyOutput{}
		resp.Body.Status = http.StatusOK
		resp.Body.Message = "User has been successfully verified!"
		return resp, nil
	})

	huma.Register(api, huma.Operation{
		OperationID:   "login-user",
		Method:        http.MethodPost,
		Path:          "/api/v1/auth/login",
		Summary:       "Logs in a user",
		Description:   "Logs in a user with their details",
		Tags:          []string{"Login"},
		DefaultStatus: http.StatusOK,
	}, func(ctx context.Context, input *LoginUserInput) (*LoginUserOutput, error) {
		user, err := authService.LoginUser(ctx, LoginUserRequest{
			Email:     input.Body.Email,
			Password:  input.Body.Password,
			UserAgent: input.Body.UserAgent,
			IPAddress: input.Body.IPAddress,
		})
		if err != nil {
			return nil, huma.Error400BadRequest(err.Error(), err)
		}

		resp := &LoginUserOutput{}
		resp.Body.Status = http.StatusOK
		resp.Body.User = user.User
		resp.Body.AccessToken = user.AccessToken
		resp.Body.RefreshToken = user.RefreshToken
		resp.Body.Message = "User successfully logged in!"

		return resp, nil
	})

	huma.Register(api, huma.Operation{
		OperationID:   "forgot-password",
		Method:        http.MethodPost,
		Path:          "/api/v1/auth/forgot-password",
		Summary:       "User forgets password",
		Description:   "User gets to generate a token to create a new password",
		Tags:          []string{"Forgot password"},
		DefaultStatus: http.StatusOK,
	}, func(ctx context.Context, input *ForgotPasswordInput) (*ForgotPasswordOutput, error) {
		err := authService.ForgotPassword(ctx, input.Body.Email)
		if err != nil {
			return nil, huma.Error400BadRequest(err.Error(), err)
		}

		resp := &ForgotPasswordOutput{}
		resp.Body.Status = http.StatusOK
		resp.Body.Message = "Password reset link sent successfully!"
		//Password reset successfully

		return resp, nil
	})
}
