package resend

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSendVerificationEmail(t *testing.T) {
	args := struct {
		email             string
		userID            string
		verificationToken string
		path              string
	}{
		email:             "mail@example.com",
		userID:            uuid.New().String(),
		verificationToken: "verificationToken",
		path:              "verify",
	}

	err := SendVerificationEmail(args.email, args.userID, args.verificationToken, args.path)
	assert.NoError(t, err)
}
