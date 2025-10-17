package service

import (
	"errors"
	"libapp/internal/models"
	"libapp/internal/repository"

	"github.com/google/uuid"
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
	if book.BookName == "" {
		return nil, errors.New("BookName cannot be empty")
	}

	book.BookID = uuid.New().String()
	//call repository
	createdBook, err := s.repo.Create(book)
	if err != nil {
		return nil, err
	}
	return createdBook, nil
}

// getallbooks method
func (s *BookService) GetAllBooks() ([]*models.Book, error) {
	booksList, err := s.repo.GetAllBooks()
	if err != nil {
		return nil, err
	}
	return booksList, nil
}
