package token

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"huma-auth/pkg/utils"
	"testing"
	"time"
)

func TestPasteoMaker(t *testing.T) {
	maker, err := NewPasetoMaker(utils.RandomString(32))
	require.NoError(t, err)

	userId := uuid.New()
	role := utils.RandomString(10)
	duration := 15 * time.Minute

	issuedAt := time.Now()
	expiredAt := issuedAt.Add(duration)

	token, err := maker.CreateToken(userId.String(), role, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)

	require.NotZero(t, payload.UserId)
	require.Equal(t, userId, uuid.MustParse(payload.UserId))
	require.Equal(t, role, payload.Role)
	require.WithinDuration(t, issuedAt, payload.IssuedAt, duration)
	require.WithinDuration(t, expiredAt, payload.ExpiredAt, duration)
}
