package repositories

import (
	"database/sql"
	"savegen-api/models"
)

type BookRepository struct {
	db *sql.DB
}

func NewBookRepository(db *sql.DB) *BookRepository {
	return &BookRepository{db: db}
}

func (r *BookRepository) GetAllBooks() ([]models.Book, error) {
	rows, err := r.db.Query("SELECT id, title, author, price FROM books ORDER BY id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []models.Book
	for rows.Next() {
		var book models.Book
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Price)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	return books, nil
} 