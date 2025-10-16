package service

import (
	"errors"
	"libapp/internal/models"
	"libapp/internal/repository"
)

// dependency injection
type BookService struct {
	repo repository.BookRepository
}

// constructor
func NewBookService(repo repository.BookRepository) *BookService {
	return &BookService{
		repo: repo,
	}
}

// create method
func (s *BookService) CreateBook(book *models.Book) (*models.Book, error) {
	//validations
	if book.BookID == 0 {
		return nil, errors.New("BookID cannot be empty")
	}
	if book.BookName == "" {
		return nil, errors.New("BookName cannot be empty")
	}

	//call repository
	createdBook, err := s.repo.Create(book)
	if err != nil {
		return nil, err
	}
	return createdBook, nil
}
