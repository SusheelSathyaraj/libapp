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

func (r *MemoryBookRepo) GetAllBooks() ([]*models.Book, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	books := make([]*models.Book, 0, len(r.books))
	for _, book := range r.books {
		books = append(books, book)
	}

	return books, nil
}

func (r *MemoryBookRepo) GetBookByID(book_id string) (*models.Book, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	book, exists := r.books[book_id]
	if !exists {
		return nil, errors.New("book not found")
	}
	return book, nil
}

func (r *MemoryBookRepo) UpdateBook(id string, updatedbook *models.Book) (*models.Book, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	book, exists := r.books[id]
	if !exists {
		return nil, errors.New("book does not exist")
	}

	if updatedbook.BookName != "" {
		book.BookName = updatedbook.BookName
	}
	if updatedbook.Author != "" {
		book.Author = updatedbook.Author
	}
	if updatedbook.Genre != "" {
		book.Genre = updatedbook.Genre
	}
	if updatedbook.ISBN != 0 {
		book.ISBN = updatedbook.ISBN
	}
	book.Available = updatedbook.Available

	r.books[id] = book

	return book, nil

}

func (r *MemoryBookRepo) DeleteBook(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, exists := r.books[id]
	if !exists {
		return errors.New("book does not exist")
	}

	delete(r.books, id)
	return nil
}
