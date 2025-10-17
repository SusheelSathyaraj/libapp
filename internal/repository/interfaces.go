package repository

import "libapp/internal/models"

type BookRepository interface {
	GetAllBooks() ([]*models.Book, error)
	//GetBookByID(book_id int) (*models.Book, error)
	Create(book *models.Book) (*models.Book, error)
}
