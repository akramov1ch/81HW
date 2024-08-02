package repository

import "81HW/internal/model"

type BookRepository interface {
	GetBooks() ([]model.Book, error)
	GetBook(id int) (*model.Book, error)
	CreateBook(book *model.Book) error
	UpdateBook(book *model.Book) error
	DeleteBook(id int) error
}