package authors

import (
	"context"
	"database/sql"

	"github.com/vertiavo/bookish-bliss-api/pkg/models"
)

// Repository is the interface that defines the methods that a repository should implement
type Repository interface {
	GetAll(ctx context.Context) ([]models.Author, error)
	GetByID(ctx context.Context, id int) (*models.Author, error)
	Create(ctx context.Context, author *models.Author) error
	Update(ctx context.Context, author *models.Author) error
	Delete(ctx context.Context, id int) error
}

type repository struct {
	db *sql.DB
}

func (r *repository) GetAll(ctx context.Context) ([]models.Author, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT id, first_name, last_name FROM authors")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	authors := []models.Author{}
	for rows.Next() {
		var author models.Author
		if err := rows.Scan(&author.ID, &author.FirstName, &author.LastName); err != nil {
			return nil, err
		}
		authors = append(authors, author)
	}
	return authors, nil
}

func (r *repository) GetByID(ctx context.Context, id int) (*models.Author, error) {
	var author models.Author
	err := r.db.QueryRowContext(ctx, "SELECT id, first_name, last_name FROM authors WHERE id = ?", id).
		Scan(&author.ID, &author.FirstName, &author.LastName)
	if err != nil {
		return nil, err
	}

	return &author, nil
}

func (r *repository) Create(ctx context.Context, author *models.Author) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO authors (first_name, last_name) VALUES (?, ?)",
		author.FirstName, author.LastName)
	return err
}

func (r *repository) Update(ctx context.Context, author *models.Author) error {
	_, err := r.db.ExecContext(ctx, "UPDATE authors SET first_name = ?, last_name = ? WHERE id = ?",
		author.FirstName, author.LastName, author.ID)
	return err
}

func (r *repository) Delete(ctx context.Context, id int) error {
	_, err := r.db.ExecContext(ctx, "DELETE FROM authors WHERE id = ?", id)
	return err
}

// NewRepository creates a new author repository
func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}
