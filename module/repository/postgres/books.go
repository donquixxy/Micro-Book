package postgres

import (
	"context"
	"errors"
	"micro-book/domain/entity"
	"micro-book/module/repository"

	"gorm.io/gorm"
)

type booksRepository struct {
	db *gorm.DB
}

func NewBooksRepository(db *gorm.DB) repository.BooksRepository {
	return &booksRepository{
		db: db,
	}
}

func (b *booksRepository) GetByID(ctx context.Context, id string) (*entity.Books, error) {
	var books *entity.Books

	err := b.db.Where("id = ?", id).First(&books).Error

	if err != nil {
		return nil, err
	}

	return books, nil
}
func (b *booksRepository) Create(ctx context.Context, value *entity.Books) error {
	result := b.db.Create(&value)

	return result.Error
}
func (b *booksRepository) GetAll(ctx context.Context) ([]*entity.Books, error) {
	var books []*entity.Books

	b.db.Find(&books).Order("created_at DESC")

	if len(books) == 0 {
		return nil, errors.New("books not found")
	}

	return books, nil
}
func (b *booksRepository) Update(ctx context.Context, newVal *entity.Books) error {
	err := b.db.Updates(&newVal).Error

	if err != nil {
		return err
	}

	return nil
}
func (b *booksRepository) GetByIDCategory(ctx context.Context, idCategory string) ([]*entity.Books, error) {
	var books []*entity.Books

	b.db.Where("id_category = ?", idCategory).Find(&books).Order("created_at DESC")

	if len(books) == 0 {
		return nil, errors.New("no books found")
	}

	return books, nil
}
