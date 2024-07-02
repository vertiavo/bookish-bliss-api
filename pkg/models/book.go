package models

// Book represents a book information
type Book struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	AuthorID int    `json:"author_id"`
	GenreID  int    `json:"genre_id"`
	Year     int    `json:"year"`
}
