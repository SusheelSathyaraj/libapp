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

// getbookbyid method
func (s *BookService) GetBookByID(book_id string) (*models.Book, error) {
	//validation
	if book_id == "" {
		return nil, errors.New("id cannot be empty")
	}
	bookDetails, err := s.repo.GetBookByID(book_id)
	if err != nil {
		return nil, err
	}
	return bookDetails, nil
}

// updateBook method
func (s *BookService) UpdateBook(book_id string, updatedBook *models.Book) (*models.Book, error) {
	//validation
	if book_id == "" {
		return nil, errors.New("bookid cannot be empty")
	}
	updatedBookDetails, err := s.repo.UpdateBook(book_id, updatedBook)
	if err != nil {
		return nil, err
	}
	return updatedBookDetails, nil
}

// deletebook method
func (s *BookService) DeleteBook(book_id string) error {
	//validation
	if book_id == "" {
		return errors.New("bookid cannot be empty")
	}
	err := s.repo.DeleteBook(book_id)
	if err != nil {
		return err
	}
	return nil
}
