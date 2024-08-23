package auth

type RegisterOutput struct {
	Body struct {
		User    UserResponse `json:"user"`
		Status  int          `json:"status"`
		Message string       `json:"message" example:"Successfully registered" doc:"Sends user message!"`
	}
}

type RegisterInput struct {
	Body UserRequest
}

type VerifyInput struct {
	UserId string `query:"user_id" doc:"Get user ID from the query param" example:"yy872y2-sbhbsd-eryjndsa-2378y4"`
	Token  string `query:"token" doc:"Get token from the query param" example:"nbdkjdnddu908ue2jnsjdn98u3kjnkndnd"`
}

type VerifyOutput struct {
	Body struct {
		Status  int `json:"status"`
		Message string
	}
}

type LoginUserInput struct {
	Body struct {
		Email     string `json:"email" doc:"User email to login" example:"user@example.com"`
		Password  string `json:"password" doc:"User password" example:"password"`
		UserAgent string `json:"user_agent,omitempty" doc:"User user agent" example:"user-agent"`
		IPAddress string `json:"ip_address,omitempty" doc:"User IP address" example:"127.0.0.1"`
	}
}

type LoginUserOutput struct {
	Body struct {
		AccessToken  string       `json:"access_token"`
		RefreshToken string       `json:"refresh_token"`
		User         UserResponse `json:"user"`
		Status       int          `json:"status"`
		Message      string       `json:"message"`
	}
}

type ForgotPasswordInput struct {
	Body struct {
		Email string `json:"email" doc:"User email to forgot" example:"user@example.com"`
	}
}

type ForgotPasswordOutput struct {
	Body struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}
}
