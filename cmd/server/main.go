package main

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	_ "github.com/danielgtaylor/huma/v2/formats/cbor"
	"github.com/danielgtaylor/huma/v2/humacli"
	"github.com/go-chi/chi/v5"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"huma-auth/config"
	"huma-auth/internal/auth"
	"huma-auth/pkg/database"
	"log"
	"net/http"
)

type Options struct {
	Port int `help:"Port to listen on" short:"p" default:"8080"`
}

func main() {
	// env config
	env, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal(err)
	}

	dbase, err := database.Connect(env.DatabaseUrl)
	if err != nil {
		log.Fatal(err)
	}

	// Close the dbase connection
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(dbase)

	err = runMigrations(env.DatabaseUrl)
	if err != nil {
		log.Fatalf("Could not run migrations: %v", err)
	}

	c := huma.DefaultConfig("Auth API", "1.0.0")
	c.Components.SecuritySchemes = map[string]*huma.SecurityScheme{
		"bearer": {
			Type:         "http",
			Scheme:       "bearer",
			BearerFormat: "JWT",
		},
	}

	// Create a new router & API
	cli := humacli.New(func(hooks humacli.Hooks, options *Options) {
		router := chi.NewMux()
		api := humachi.New(router, c)

		// Register Routes
		auth.RegisterHandlers(api)

		hooks.OnStart(func() {
			fmt.Printf("Starting server on port %d...\n", options.Port)
			err := http.ListenAndServe(fmt.Sprintf(":%d", options.Port), router)
			if err != nil {
				log.Fatal(err)
			}
		})
	})

	// Start the server!
	cli.Run()
}

// runMigrations applies all the migrations from the specified directory
func runMigrations(dataSourceName string) error {
	// Connect to the database for migration
	dbase, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return err
	}
	defer func(dbase *sql.DB) {
		err := dbase.Close()
		if err != nil {

		}
	}(dbase)

	// Set up file-based source driver
	sourceURL := "file://migrations"
	sourceDriver, err := (&file.File{}).Open(sourceURL)
	if err != nil {
		return err
	}

	// Set up the PostgreSQL database driver
	dbDriver, err := postgres.WithInstance(dbase, &postgres.Config{})
	if err != nil {
		return err
	}

	// Create a new migrate instance
	m, err := migrate.NewWithInstance("file", sourceDriver, "postgres", dbDriver)
	if err != nil {
		return err
	}

	// Run migrations
	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	log.Println("Migrations ran successfully.")
	return nil
}
