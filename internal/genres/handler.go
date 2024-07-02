package genres

import (
	"database/sql"
	"strconv"

	"github.com/gofiber/fiber/v3"
	appcontext "github.com/vertiavo/bookish-bliss-api/internal/context"
	"github.com/vertiavo/bookish-bliss-api/pkg/models"
)

func RegisterHandlers(app *fiber.App, db *sql.DB) {
	s := NewService(NewRepository(db))
	h := NewHandler(s)

	app.Get("/genres", h.GetAll)
	app.Get("/genres/:id", h.GetByID)
	app.Post("/genres", h.Create)
	app.Put("/genres", h.Update)
	app.Delete("/genres/:id", h.Delete)
}

type Handler interface {
	GetAll(c fiber.Ctx) error
	GetByID(c fiber.Ctx) error
	Create(c fiber.Ctx) error
	Update(c fiber.Ctx) error
	Delete(c fiber.Ctx) error
}

type handler struct {
	service Service
}

func (h *handler) GetAll(c fiber.Ctx) error {
	ctx := c.UserContext()
	logger := appcontext.GetLogger(ctx)

	authors, err := h.service.GetAll(ctx)
	if err != nil {
		logger.Println("Error fetching genres:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(authors)
}

func (h *handler) GetByID(c fiber.Ctx) error {
	ctx := c.UserContext()
	logger := appcontext.GetLogger(ctx)

	id := c.Params("id")
	genreId, err := strconv.Atoi(id)
	if err != nil {
		logger.Println("Error parsing genre ID:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid author ID",
		})
	}

	author, err := h.service.GetByID(ctx, genreId)
	if err != nil {
		logger.Println("Error fetching genre by ID:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(author)
}

func (h *handler) Create(c fiber.Ctx) error {
	ctx := c.UserContext()
	logger := appcontext.GetLogger(ctx)

	var genre models.Genre
	if err := c.Bind().Body(&genre); err != nil {
		logger.Println("Error binding genre:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := h.service.Create(ctx, &genre); err != nil {
		logger.Println("Error creating genre:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(genre)
}

func (h *handler) Update(c fiber.Ctx) error {
	ctx := c.UserContext()
	logger := appcontext.GetLogger(ctx)

	var genre models.Genre
	if err := c.Bind().Body(&genre); err != nil {
		logger.Println("Error binding genre:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := h.service.Update(ctx, &genre); err != nil {
		logger.Println("Error updating genre:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(genre)
}

func (h *handler) Delete(c fiber.Ctx) error {
	ctx := c.UserContext()
	logger := appcontext.GetLogger(ctx)

	id := c.Params("id")
	genreId, err := strconv.Atoi(id)
	if err != nil {
		logger.Println("Error parsing genre ID:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid genre ID",
		})
	}

	if err = h.service.Delete(ctx, genreId); err != nil {
		logger.Println("Error deleting genre:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// NewHandler creates a new genre handler with the provided service
func NewHandler(s Service) Handler {
	return &handler{
		service: s,
	}
}
