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
