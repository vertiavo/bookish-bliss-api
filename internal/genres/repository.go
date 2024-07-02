package genres

import (
	"context"
	"database/sql"

	"github.com/vertiavo/bookish-bliss-api/pkg/models"
)

type Repository interface {
	GetAll(ctx context.Context) ([]models.Genre, error)
	GetByID(ctx context.Context, id int) (*models.Genre, error)
	Create(ctx context.Context, genre *models.Genre) error
	Update(ctx context.Context, genre *models.Genre) error
	Delete(ctx context.Context, id int) error
}

type repository struct {
	db *sql.DB
}

func (r *repository) GetAll(ctx context.Context) ([]models.Genre, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT id, name FROM genres")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	genres := []models.Genre{}
	for rows.Next() {
		var genre models.Genre
		if err := rows.Scan(&genre.ID, &genre.Name); err != nil {
			return nil, err
		}
		genres = append(genres, genre)
	}
	return genres, nil
}

func (r *repository) GetByID(ctx context.Context, id int) (*models.Genre, error) {
	var genre models.Genre
	err := r.db.QueryRowContext(ctx, "SELECT id, name FROM genres WHERE id = ?", id).
		Scan(&genre.ID, &genre.Name)
	if err != nil {
		return nil, err
	}

	return &genre, nil
}

func (r *repository) Create(ctx context.Context, genre *models.Genre) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO genres (name) VALUES (?)",
		genre.Name)
	return err
}

func (r *repository) Update(ctx context.Context, genre *models.Genre) error {
	_, err := r.db.ExecContext(ctx, "UPDATE genres SET name = ? WHERE id = ?",
		genre.Name, genre.ID)
	return err
}

func (r *repository) Delete(ctx context.Context, id int) error {
	_, err := r.db.ExecContext(ctx, "DELETE FROM genres WHERE id = ?", id)
	return err
}

// NewRepository creates a new genre repository
func NewRepository(db *sql.DB) Repository {
	return &repository{db}
}
