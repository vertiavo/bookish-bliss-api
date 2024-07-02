package app

import (
	"context"
	"log"
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/vertiavo/bookish-bliss-api/internal/authors"
	"github.com/vertiavo/bookish-bliss-api/internal/books"
	"github.com/vertiavo/bookish-bliss-api/internal/config"
	appcontext "github.com/vertiavo/bookish-bliss-api/internal/context"
	"github.com/vertiavo/bookish-bliss-api/internal/database"
	"github.com/vertiavo/bookish-bliss-api/internal/genres"
)

// Initialize sets up the application
func Initialize() {
	// Load configuration
	cfg := config.LoadConfig()

	// Set up logger
	logger := log.New(os.Stdout, "", log.LstdFlags)

	// Create a base context
	ctx := context.Background()
	ctx = appcontext.WithConfig(ctx, cfg)
	ctx = appcontext.WithLogger(ctx, logger)

	db, err := database.InitDB(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Create a new Fiber app
	app := fiber.New()

	app.Use(func(c fiber.Ctx) error {
		c.SetUserContext(ctx)
		return c.Next()
	})

	// Register routes
	authors.RegisterHandlers(app, db)
	books.RegisterHandlers(app, db)
	genres.RegisterHandlers(app, db)

	log.Fatal(app.Listen(":3000"))
}
