package repository

import (
	"database/sql"
	"81HW/internal/model"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{db: db}
}

func (r *PostgresRepository) GetBooks() ([]model.Book, error) {
	m := []model.Book{}
	rows, err := r.db.Query("SELECT * FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var b model.Book
		err := rows.Scan(&b.ID, &b.Title, &b.Author, &b.PublishedYear)
		if err != nil {
			return nil, err
		}
		m = append(m, b)
	}
	return m, nil
}

func (r *PostgresRepository) GetBook(id int) (*model.Book, error) {
	var b model.Book
	err := r.db.QueryRow("SELECT * FROM books WHERE id = $1", id).Scan(&b.ID, &b.Title, &b.Author, &b.PublishedYear)
	if err != nil {
		return nil, err
	}
	return &b, nil
}

func (r *PostgresRepository) CreateBook(book *model.Book) error {
	_, err := r.db.Exec("INSERT INTO books (title, author, published_year) VALUES ($1, $2, $3)", book.Title, book.Author, book.PublishedYear)
	return err
}

func (r *PostgresRepository) UpdateBook(book *model.Book) error {
	_, err := r.db.Exec("UPDATE books SET title = $1, author = $2, published_year = $3 WHERE id = $4", book.Title, book.Author, book.PublishedYear, book.ID)
	return err
}

func (r *PostgresRepository) DeleteBook(id int) error {
	_, err := r.db.Exec("DELETE FROM books WHERE id = $1", id)
	return err
}

// GetBooks, GetBook, CreateBook, UpdateBook, DeleteBook metodlarini shu yerda amalga oshiring