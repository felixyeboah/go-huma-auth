package auth

import (
	"context"
	"github.com/danielgtaylor/huma/v2/humatest"
	"testing"
)

var ctx = context.Background()

func TestCreateUser(t *testing.T) {
	_, api := humatest.New(t)

	RegisterHandlers(api)

	resp := api.PostCtx(ctx, "/api/v1/auth/register", map[string]string{
		"name":         "Test Subject",
		"email":        "mail@example.com",
		"phone_number": "2568954252",
		"password":     "secure-password",
		"role_name":    "admin",
	})

	if resp.Code != 201 {
		t.Fatalf("Unexpected status code: %d", resp.Code)
	}
}

func TestCreateUserError(t *testing.T) {
	_, api := humatest.New(t)

	RegisterHandlers(api)

	resp := api.PostCtx(ctx, "/api/v1/auth/register", map[string]string{
		"name":         "Test Subject",
		"email":        "mail@example.com",
		"phone_number": "2568954252",
	})

	if resp.Code != 422 {
		t.Fatalf("Unexpected status code: %d", resp.Code)
	}
}

func TestVerifyUser(t *testing.T) {
	_, api := humatest.New(t)

	RegisterHandlers(api)

	resp := api.Get("/api/v1/auth/verify?user_id=786246246814-4tbxdhdx&token=46783467464njdjasdnhndjd8u498324kjndshd")
	if resp.Code != 200 {
		t.Fatalf("Unexpected status code: %d", resp.Code)
	}
}
