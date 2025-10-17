package memory

import (
	"errors"
	"libapp/internal/models"
	"sync"
)

type MemoryBookRepo struct {
	mu    sync.Mutex
	books map[string]*models.Book
}

func NewMemoryBookRepo() *MemoryBookRepo {
	return &MemoryBookRepo{
		books: make(map[string]*models.Book),
	}
}

func (r *MemoryBookRepo) Create(book *models.Book) (*models.Book, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, exists := r.books[book.BookID]; exists {
		return nil, errors.New("book already exists")
	}
	//add the book to the map
	r.books[book.BookID] = book
	return book, nil
}
