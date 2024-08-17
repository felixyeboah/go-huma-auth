// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Privilege struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Role struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type RolePrivilege struct {
	RoleID      uuid.UUID `json:"role_id"`
	PrivilegeID uuid.UUID `json:"privilege_id"`
}

type Session struct {
	ID             uuid.UUID      `json:"id"`
	AccessToken    string         `json:"access_token"`
	RefreshToken   string         `json:"refresh_token"`
	UserID         uuid.UUID      `json:"user_id"`
	ExpiryDate     time.Time      `json:"expiry_date"`
	UserAgent      sql.NullString `json:"user_agent"`
	IpAddress      sql.NullString `json:"ip_address"`
	LastAccessedAt sql.NullTime   `json:"last_accessed_at"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
}

type User struct {
	ID          uuid.UUID      `json:"id"`
	Name        string         `json:"name"`
	Email       string         `json:"email"`
	Avatar      sql.NullString `json:"avatar"`
	PhoneNumber string         `json:"phone_number"`
	Password    string         `json:"password"`
	IsVerified  bool           `json:"is_verified"`
	RoleID      uuid.UUID      `json:"role_id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
}
