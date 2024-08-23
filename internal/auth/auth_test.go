package auth

import (
	"context"
	"database/sql"
	"github.com/danielgtaylor/huma/v2/humatest"
	"huma-auth/pkg/database"
	"log"
	"testing"
)

var ctx = context.Background()

func TestCreateUser(t *testing.T) {
	_, api := humatest.New(t)

	dbase, err := database.Connect()
	if err != nil {
		log.Fatalln(err)
	}

	// Close the dbase connection
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(dbase)

	RegisterHandlers(api, dbase)

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

	dbase, err := database.Connect()
	if err != nil {
		log.Fatalln(err)
	}

	// Close the dbase connection
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(dbase)

	RegisterHandlers(api, dbase)

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

	dbase, err := database.Connect()
	if err != nil {
		log.Fatalln(err)
	}

	// Close the dbase connection
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(dbase)

	RegisterHandlers(api, dbase)

	resp := api.Get("/api/v1/auth/verify?user_id=786246246814-4tbxdhdx&token=46783467464njdjasdnhndjd8u498324kjndshd")
	if resp.Code != 200 {
		t.Fatalf("Unexpected status code: %d", resp.Code)
	}
}

func TestVerifyUserError(t *testing.T) {
	_, api := humatest.New(t)

	dbase, err := database.Connect()
	if err != nil {
		log.Fatalln(err)
	}

	// Close the dbase connection
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(dbase)

	RegisterHandlers(api, dbase)

	resp := api.Get("/api/v1/auth/verify")
	if resp.Code != 422 {
		t.Fatalf("Unexpected status code: %d", resp.Code)
	}
}

func TestLogin(t *testing.T) {
	_, api := humatest.New(t)

	dbase, err := database.Connect()
	if err != nil {
		log.Fatalln(err)
	}

	// Close the dbase connection
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(dbase)

	RegisterHandlers(api, dbase)

	resp := api.PostCtx(ctx, "/api/v1/auth/login", map[string]string{
		"email":    "mail@example.com",
		"password": "secure-password",
	})
	if resp.Code != 200 {
		t.Fatalf("Unexpected status code: %d", resp.Code)
	}
}
