package books

import (
	"context"

	"github.com/vertiavo/bookish-bliss-api/pkg/models"
)

// Service represents the book service
type Service interface {
	GetAll(ctx context.Context) ([]models.Book, error)
	GetByID(ctx context.Context, id int) (*models.Book, error)
	Create(ctx context.Context, book *models.Book) error
	Update(ctx context.Context, book *models.Book) error
	Delete(ctx context.Context, id int) error
}

type service struct {
	repo Repository
}

func (s *service) GetAll(ctx context.Context) ([]models.Book, error) {
	return s.repo.GetAll(ctx)
}

func (s *service) GetByID(ctx context.Context, id int) (*models.Book, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *service) Create(ctx context.Context, book *models.Book) error {
	return s.repo.Create(ctx, book)
}

func (s *service) Update(ctx context.Context, book *models.Book) error {
	return s.repo.Update(ctx, book)
}

func (s *service) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}

// NewService creates a new book service
func NewService(repo Repository) Service {
	return &service{repo: repo}
}
