package auth

type RegisterOutput struct {
	Body UserResponse
}

type RegisterInput struct {
	Body UserRequest
}
