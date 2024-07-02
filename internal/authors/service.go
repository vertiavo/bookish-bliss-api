package authors

import (
	"context"

	"github.com/vertiavo/bookish-bliss-api/pkg/models"
)

// Service represents the author service
type Service interface {
	GetAll(ctx context.Context) ([]models.Author, error)
	GetByID(ctx context.Context, id int) (*models.Author, error)
	Create(ctx context.Context, author *models.Author) error
	Update(ctx context.Context, author *models.Author) error
	Delete(ctx context.Context, id int) error
}

type service struct {
	repo Repository
}

func (s *service) GetAll(ctx context.Context) ([]models.Author, error) {
	return s.repo.GetAll(ctx)
}

func (s *service) GetByID(ctx context.Context, id int) (*models.Author, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *service) Create(ctx context.Context, author *models.Author) error {
	return s.repo.Create(ctx, author)
}

func (s *service) Update(ctx context.Context, author *models.Author) error {
	return s.repo.Update(ctx, author)
}

func (s *service) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}

// NewService creates a new author service
func NewService(repo Repository) Service {
	return &service{repo: repo}
}
