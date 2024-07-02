package books

import (
	"context"
	"database/sql"

	"github.com/vertiavo/bookish-bliss-api/pkg/models"
)

type Repository interface {
	GetAll(ctx context.Context) ([]models.Book, error)
	GetByID(ctx context.Context, id int) (*models.Book, error)
	Create(ctx context.Context, book *models.Book) error
	Update(ctx context.Context, book *models.Book) error
	Delete(ctx context.Context, id int) error
}

type repository struct {
	db *sql.DB
}

func (r *repository) GetAll(ctx context.Context) ([]models.Book, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT id, title, author_id, genre_id, year FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	books := []models.Book{}
	for rows.Next() {
		var book models.Book
		if err := rows.Scan(&book.ID, &book.Title, &book.AuthorID, &book.GenreID, &book.Year); err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}

func (r *repository) GetByID(ctx context.Context, id int) (*models.Book, error) {
	var book models.Book
	err := r.db.QueryRowContext(ctx, "SELECT id, title, author_id, genre_id, year FROM books WHERE id = ?", id).
		Scan(&book.ID, &book.Title, &book.AuthorID, &book.GenreID, &book.Year)
	if err != nil {
		return nil, err
	}

	return &book, nil
}

func (r *repository) Create(ctx context.Context, book *models.Book) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO books (title, author_id, genre_id, year) VALUES (?, ?, ?, ?)",
		book.Title, book.AuthorID, book.GenreID, book.Year)
	return err
}

func (r *repository) Update(ctx context.Context, book *models.Book) error {
	_, err := r.db.ExecContext(ctx, "UPDATE books SET title = ?, author_id = ?, genre_id = ?, year = ? WHERE id = ?",
		book.Title, book.AuthorID, book.GenreID, book.Year, book.ID)
	return err
}

func (r *repository) Delete(ctx context.Context, id int) error {
	_, err := r.db.ExecContext(ctx, "DELETE FROM books WHERE id = ?", id)
	return err
}

// NewRepository creates a new book repository
func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}
