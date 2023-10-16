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

	err := b.db.WithContext(ctx).Where("id = ?", id).First(&books).
		Preload("Category").Preload("Genre").
		Error

	if err != nil {
		return nil, err
	}

	return books, nil
}
func (b *booksRepository) Create(ctx context.Context, value *entity.Books) error {
	result := b.db.WithContext(ctx).Create(&value)

	return result.Error
}
func (b *booksRepository) GetAll(ctx context.Context) ([]*entity.Books, error) {
	var books []*entity.Books

	b.db.WithContext(ctx).Order("created_at DESC").Preload("Category").Preload("Genre").Find(&books)

	if len(books) == 0 {
		return nil, errors.New("books not found")
	}

	return books, nil
}
func (b *booksRepository) Update(ctx context.Context, newVal *entity.Books) error {
	err := b.db.WithContext(ctx).Updates(&newVal).Error

	if err != nil {
		return err
	}

	return nil
}
func (b *booksRepository) GetByIDCategory(ctx context.Context, idCategory string) ([]*entity.Books, error) {
	var books []*entity.Books

	b.db.WithContext(ctx).Where("id_category = ?", idCategory).Find(&books).Order("created_at DESC").
		Preload("Category").Preload("Genre")

	if len(books) == 0 {
		return nil, errors.New("no books found")
	}

	return books, nil
}

func (b *booksRepository) CreateCategory(ctx context.Context, value *entity.Category) error {
	err := b.db.WithContext(ctx).Create(&value).Error

	return err
}
func (b *booksRepository) CreateGenre(ctx context.Context, value *entity.Genre) error {
	err := b.db.WithContext(ctx).Create(&value).Error

	return err
}
