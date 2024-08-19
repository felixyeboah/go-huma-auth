package token

import (
	"time"
)

type Payload struct {
	UserId    string    `json:"userId"`
	Role      string    `json:"role"`
	ExpiredAt time.Time `json:"expiredAt"`
	IssuedAt  time.Time `json:"issuedAt"`
}

func NewPayload(UserId, Role string, duration time.Duration) (*Payload, error) {
	now := time.Now()
	return &Payload{
		UserId:    UserId,
		Role:      Role,
		IssuedAt:  now,
		ExpiredAt: now.Add(duration),
	}, nil
}
