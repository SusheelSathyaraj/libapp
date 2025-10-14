package memory

import (
	"errors"
	"libapp/internal/models"
)

type MemoryBookRepo struct {
	books map[int]*models.Book
}

func NewMemoryBookRepo() *MemoryBookRepo {
	return &MemoryBookRepo{
		books: make(map[int]*models.Book),
	}
}

func (r *MemoryBookRepo) Create(book *models.Book) error {
	if _, exists := r.books[book.BookID]; exists {
		return errors.New("book already exists")
	}
	//add the book to the map
	r.books[book.BookID] = book
	return nil
}
