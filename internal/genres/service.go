package genres

import (
	"context"

	"github.com/vertiavo/bookish-bliss-api/pkg/models"
)

// Service represents the genre service
type Service interface {
	GetAll(ctx context.Context) ([]models.Genre, error)
	GetByID(ctx context.Context, id int) (*models.Genre, error)
	Create(ctx context.Context, genre *models.Genre) error
	Update(ctx context.Context, genre *models.Genre) error
	Delete(ctx context.Context, id int) error
}

type service struct {
	repo Repository
}

func (s *service) GetAll(ctx context.Context) ([]models.Genre, error) {
	return s.repo.GetAll(ctx)
}

func (s *service) GetByID(ctx context.Context, id int) (*models.Genre, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *service) Create(ctx context.Context, genre *models.Genre) error {
	return s.repo.Create(ctx, genre)
}

func (s *service) Update(ctx context.Context, genre *models.Genre) error {
	return s.repo.Update(ctx, genre)
}

func (s *service) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}

// NewService creates a new genre service
func NewService(repo Repository) Service {
	return &service{repo: repo}
}
