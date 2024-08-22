package auth

import (
	"github.com/google/uuid"
	"time"
)

type UserRequest struct {
	Name        string `json:"name" maxLength:"50" example:"John Doe" doc:"Register a user's name" required:"true"`
	Email       string `json:"email" example:"johndoe@example.com" doc:"Register a user's email" required:"true"`
	PhoneNumber string `json:"phone_number" minimum:"10" example:"1381234567" doc:"Register a user's phone number" required:"true"`
	Password    string `json:"password" minimum:"8" example:"123456" doc:"Register a user's password" required:"true"`
	RoleName    string `json:"role_name" example:"admin" doc:"Register a user's role"  required:"true" enum:"user,admin"`
}

type UserResponse struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	IsVerified  bool      `json:"is_verified"`
	RoleID      uuid.UUID `json:"role_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
