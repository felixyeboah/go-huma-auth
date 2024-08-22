package auth

import (
	"github.com/danielgtaylor/huma/v2/humatest"
	"testing"
)

func TestCreateUser(t *testing.T) {
	_, api := humatest.New(t)

	RegisterHandlers(api)

	resp := api.Post("/api/v1/auth/register", map[string]string{
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

	resp := api.Post("/api/v1/auth/register", map[string]string{
		"name":         "Test Subject",
		"email":        "mail@example.com",
		"phone_number": "2568954252",
	})

	if resp.Code != 422 {
		t.Fatalf("Unexpected status code: %d", resp.Code)
	}
}
