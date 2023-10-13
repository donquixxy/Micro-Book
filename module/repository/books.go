package repository

import (
	"context"
	"micro-book/domain/entity"
)

// Booksrepository is a contract for booksrepository
// booksrepository can be used with any database connection
type BooksRepository interface {
	GetByID(ctx context.Context, id string) (*entity.Books, error)
	Create(ctx context.Context, value *entity.Books) error
	GetAll(ctx context.Context) ([]*entity.Books, error)
	Update(ctx context.Context, newVal *entity.Books) error
	GetByIDCategory(ctx context.Context, idCategory string) ([]*entity.Books, error)
	CreateCategory(ctx context.Context, value *entity.Category) error
	CreateGenre(ctx context.Context, value *entity.Genre) error
}
