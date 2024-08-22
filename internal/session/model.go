package session

import (
	"github.com/google/uuid"
	"time"
)

type Payload struct {
	UserId    uuid.UUID `json:"userId"`
	Role      string    `json:"role"`
	UserAgent string    `json:"userAgent"`
	IPAddress string    `json:"ipAddress"`
}

type Response struct {
	ID             uuid.UUID `json:"id"`
	AccessToken    string    `json:"accessToken"`
	RefreshToken   string    `json:"refresh_token"`
	UserID         uuid.UUID `json:"user_id"`
	ExpiryDate     time.Time `json:"expiry_date"`
	UserAgent      string    `json:"user_agent"`
	IpAddress      string    `json:"ip_address"`
	LastAccessedAt time.Time `json:"last_accessed_at"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
