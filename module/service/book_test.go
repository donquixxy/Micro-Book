package service

import (
	"context"
	"micro-book/domain/entity"
	"micro-book/domain/request"
	"micro-book/utils"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockBookRepository struct {
	mock.Mock
}

func (m *MockBookRepository) GetByID(ctx context.Context, id string) (*entity.Books, error) {

	args := m.Called(ctx, id)

	return args.Get(0).(*entity.Books), args.Error(1)
}

func (m *MockBookRepository) Create(ctx context.Context, value *entity.Books) error {
	args := m.Called(ctx, value)

	return args.Error(0)
}

func (m *MockBookRepository) GetAll(ctx context.Context) ([]*entity.Books, error) {
	args := m.Called(ctx)

	return args.Get(0).([]*entity.Books), args.Error(1)
}

func (m *MockBookRepository) CreateCategory(ctx context.Context, value *entity.Category) error {
	args := m.Called(ctx, value)

	return args.Error(0)
}

func (m *MockBookRepository) CreateGenre(ctx context.Context, value *entity.Genre) error {
	args := m.Called(ctx, value)

	return args.Error(0)
}

func (m *MockBookRepository) GetByIDCategory(ctx context.Context, idCategory string) ([]*entity.Books, error) {
	args := m.Called(ctx, idCategory)

	return args.Get(0).([]*entity.Books), args.Error(1)
}

func (m *MockBookRepository) Update(ctx context.Context, newVal *entity.Books) error {
	args := m.Called(ctx, newVal)

	return args.Error(0)
}

func TestGetById(t *testing.T) {

	repo := new(MockBookRepository)

	service := NewBookService(repo)

	id := "11"

	book := &entity.Books{
		ID:         id,
		IDCategory: "123",
		IDGenre:    "123",
		Title:      "war",
		ISBN:       "921",
		Price:      100000,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
	ctx := context.Background()
	repo.On("GetByID", ctx, id).Return(book, nil).Once()
	result, err := service.GetByID(ctx, id)

	assert.NoError(t, err)
	assert.NotNil(t, result)

	assert.Equal(t, id, book.ID)
}

func TestCreateBook(t *testing.T) {
	repo := new(MockBookRepository)

	service := NewBookService(repo)

	ctx := context.Background()
	repo.On("Create", ctx, mock.Anything).Return(nil).Once()

	result, err := service.CreateBook(ctx, &request.CreateBookRequest{
		IDCategory: "11",
		IDGenre:    "123",
		Title:      "Kaguya",
		ISBN:       "9981-0021",
		Price:      80000,
	})

	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestGetAllBook(t *testing.T) {
	repo := new(MockBookRepository)

	service := NewBookService(repo)

	ctx := context.Background()

	boooks := []*entity.Books{
		GenerateRandomBooks(), GenerateRandomBooks(),
	}

	repo.On("GetAll", ctx).Return(boooks, nil).Once()

	result, err := service.GetAll(ctx)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, result, boooks)
}

func TestUpdateBook(t *testing.T) {
	repo := new(MockBookRepository)

	service := NewBookService(repo)

	ctx := context.Background()

	// Before Updated
	bef := GenerateRandomBooks()

	repo.On("Update", ctx, mock.Anything).Return(nil)

	v := &request.UpdateBookRequest{
		ID:         bef.ID,
		IDCategory: "8888",
		IDGenre:    "00000",
		Title:      "Januari",
		ISBN:       "9981-0912",
		Price:      10000,
	}

	result, err := service.Update(ctx, v)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, bef.ID, result.ID)
	assert.NotEqual(t, bef.IDCategory, result.IDCategory)
	assert.NotEqual(t, bef.IDGenre, result.IDGenre)
	assert.NotEqual(t, bef.Title, result.Title)
	assert.NotEqual(t, bef.ISBN, result.ISBN)
	assert.NotEqual(t, bef.Price, result.Price)
}

func TestGetByIDCategory(t *testing.T) {
	repo := new(MockBookRepository)

	service := NewBookService(repo)

	idCat := "9881"

	ctx := context.Background()

	listBooks := []*entity.Books{
		GenerateRandomBooks(), GenerateRandomBooks(),
	}

	repo.On("GetByIDCategory", ctx, idCat).Return(listBooks, nil).Once()

	result, err := service.GetByIDCategory(ctx, idCat)

	assert.NoError(t, err)
	assert.NotEmpty(t, result)

	assert.Equal(t, result, listBooks)
}

func TestCreateCategory(t *testing.T) {
	repo := new(MockBookRepository)

	service := NewBookService(repo)

	ctx := context.Background()

	repo.On("CreateCategory", ctx, mock.Anything).Return(nil).Once()

	result, err := service.CreateCategory(ctx, mock.Anything)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.NotEmpty(t, result.ID)
	assert.NotEmpty(t, result.Category)
	assert.NotEmpty(t, result.CreatedAt)
}

func TestCreateGenre(t *testing.T) {
	repo := new(MockBookRepository)

	service := NewBookService(repo)

	ctx := context.Background()

	repo.On("CreateGenre", ctx, mock.Anything).Return(nil).Once()

	result, err := service.CreateGenre(ctx, mock.Anything)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.NotEmpty(t, result.ID)
	assert.NotEmpty(t, result.Genre)
}

func GenerateRandomBooks() *entity.Books {
	return &entity.Books{
		ID:         utils.RandomIDGenerator(),
		IDCategory: "9881",
		IDGenre:    "221",
		Title:      "War Love",
		ISBN:       "9921-21212",
		Price:      80000,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
}
