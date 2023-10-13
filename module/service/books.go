package service

import (
	"context"
	"log"
	"micro-book/domain/entity"
	"micro-book/domain/request"
	"micro-book/module/repository"
	"micro-book/utils"
	"time"
)

type BookService interface {
	CreateBook(ctx context.Context, book *request.CreateBookRequest) (*entity.Books, error)
	CreateGenre(ctx context.Context, genre string) (*entity.Genre, error)
	CreateCategory(ctx context.Context, category string) (*entity.Category, error)
	GetByID(ctx context.Context, id string) (*entity.Books, error)
	GetAll(ctx context.Context) ([]*entity.Books, error)
	GetByIDCategory(ctx context.Context, idCategory string) ([]*entity.Books, error)
}

type bookService struct {
	bookRepository repository.BooksRepository
}

func NewBookService(
	bookRepository repository.BooksRepository,
) BookService {
	return &bookService{
		bookRepository: bookRepository,
	}
}

func (s *bookService) CreateBook(ctx context.Context, book *request.CreateBookRequest) (*entity.Books, error) {
	// Convert request to entity
	b := &entity.Books{
		ID:         utils.RandomIDGenerator(),
		IDCategory: book.IDCategory,
		IDGenre:    book.IDGenre,
		Title:      book.Title,
		ISBN:       book.ISBN,
		Price:      book.Price,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	// Insert to db
	err := s.bookRepository.Create(ctx, b)

	if err != nil {
		log.Printf("failed create books : %v", err)
		return nil, err
	}

	return b, nil
}

func (s *bookService) CreateGenre(ctx context.Context, genre string) (*entity.Genre, error) {
	gen := &entity.Genre{
		ID:        utils.RandomIDGenerator(),
		Genre:     genre,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := s.bookRepository.CreateGenre(ctx, gen)

	if err != nil {
		log.Printf("failed create genre : %v", err)
		return nil, err
	}

	return gen, nil
}

func (s *bookService) CreateCategory(ctx context.Context, category string) (*entity.Category, error) {
	cat := &entity.Category{
		ID:        utils.RandomIDGenerator(),
		Category:  category,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := s.bookRepository.CreateCategory(ctx, cat)

	if err != nil {
		log.Printf("failed create category : %v", err)
		return nil, err
	}

	return cat, nil
}

func (s *bookService) GetByID(ctx context.Context, id string) (*entity.Books, error) {
	return s.bookRepository.GetByID(ctx, id)
}

func (s *bookService) GetAll(ctx context.Context) ([]*entity.Books, error) {
	books, err := s.bookRepository.GetAll(ctx)

	if err != nil {
		return nil, err
	}

	return books, nil
}

func (s *bookService) GetByIDCategory(ctx context.Context, idCategory string) ([]*entity.Books, error) {
	books, err := s.bookRepository.GetByIDCategory(ctx, idCategory)

	if err != nil {
		return nil, err
	}

	return books, nil
}
