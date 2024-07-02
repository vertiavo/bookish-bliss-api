package authors

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

	app.Get("/authors", h.GetAll)
	app.Get("/authors/:id", h.GetByID)
	app.Post("/authors", h.Create)
	app.Put("/authors", h.Update)
	app.Delete("/authors/:id", h.Delete)
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
		logger.Println("Error fetching authors:", err)
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
	authorId, err := strconv.Atoi(id)
	if err != nil {
		logger.Println("Error parsing author ID:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid author ID",
		})
	}

	author, err := h.service.GetByID(ctx, authorId)
	if err != nil {
		logger.Println("Error fetching author by ID:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(author)
}

func (h *handler) Create(c fiber.Ctx) error {
	ctx := c.UserContext()
	logger := appcontext.GetLogger(ctx)

	var author models.Author
	if err := c.Bind().Body(&author); err != nil {
		logger.Println("Error binding author:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := h.service.Create(ctx, &author); err != nil {
		logger.Println("Error creating author:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(author)
}

func (h *handler) Update(c fiber.Ctx) error {
	ctx := c.UserContext()
	logger := appcontext.GetLogger(ctx)

	var author models.Author
	if err := c.Bind().Body(&author); err != nil {
		logger.Println("Error binding author:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := h.service.Update(ctx, &author); err != nil {
		logger.Println("Error updating author:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(author)
}

func (h *handler) Delete(c fiber.Ctx) error {
	ctx := c.UserContext()
	logger := appcontext.GetLogger(ctx)

	id := c.Params("id")
	authorId, err := strconv.Atoi(id)
	if err != nil {
		logger.Println("Error parsing author ID:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid author ID",
		})
	}

	if err = h.service.Delete(c.Context(), authorId); err != nil {
		logger.Println("Error deleting author:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// NewHandler creates a new author handler with the provided service
func NewHandler(s Service) Handler {
	return &handler{
		service: s,
	}
}
