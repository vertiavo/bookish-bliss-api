package books

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

	app.Get("/books", h.GetAll)
	app.Get("/books/:id", h.GetByID)
	app.Post("/books", h.Create)
	app.Put("/books", h.Update)
	app.Delete("/books/:id", h.Delete)
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

	books, err := h.service.GetAll(ctx)
	if err != nil {
		logger.Println("Error fetching books:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(books)
}

func (h *handler) GetByID(c fiber.Ctx) error {
	ctx := c.UserContext()
	logger := appcontext.GetLogger(ctx)

	id := c.Params("id")
	bookId, err := strconv.Atoi(id)
	if err != nil {
		logger.Println("Invalid book ID:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid book ID",
		})
	}

	book, err := h.service.GetByID(ctx, bookId)
	if err != nil {
		logger.Println("Error fetching book:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(book)
}

func (h *handler) Create(c fiber.Ctx) error {
	ctx := c.UserContext()
	logger := appcontext.GetLogger(ctx)

	var book models.Book
	if err := c.Bind().Body(&book); err != nil {
		logger.Println("Error binding book:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := h.service.Create(ctx, &book); err != nil {
		logger.Println("Error creating book:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(book)
}

func (h *handler) Update(c fiber.Ctx) error {
	ctx := c.UserContext()
	logger := appcontext.GetLogger(ctx)

	var book models.Book
	if err := c.Bind().Body(&book); err != nil {
		logger.Println("Error binding book:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	err := h.service.Update(ctx, &book)
	if err != nil {
		logger.Println("Error updating book:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(book)
}

func (h *handler) Delete(c fiber.Ctx) error {
	ctx := c.UserContext()
	logger := appcontext.GetLogger(ctx)

	id := c.Params("id")
	bookId, err := strconv.Atoi(id)
	if err != nil {
		logger.Println("Invalid book ID:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid book ID",
		})
	}

	if err = h.service.Delete(c.UserContext(), bookId); err != nil {
		logger.Println("Error deleting book:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// NewHandler creates a new book handler with the provided service
func NewHandler(s Service) Handler {
	return &handler{
		service: s,
	}
}
