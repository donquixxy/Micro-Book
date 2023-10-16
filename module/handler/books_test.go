package handler_test

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"micro-book/config"
	"micro-book/domain/entity"
	"micro-book/domain/request"
	mock_service "micro-book/mockgen"
	"micro-book/server"
	"micro-book/utils"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

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

func TestGetAll(t *testing.T) {

	ctrl := gomock.NewController(t)

	// Init testServer
	conf := config.InitConfig()
	server := server.NewServer("development")
	server.ServerWrapper(config.InitPostgres(conf))

	defer ctrl.Finish()

	service := mock_service.NewMockBookService(ctrl)

	listBook := []*entity.Books{
		GenerateRandomBooks(), GenerateRandomBooks(),
	}

	// Stubs
	service.EXPECT().GetAll(gomock.Any()).Times(1).Return(listBook, nil).AnyTimes()

	url := "/api/books"

	recorder := httptest.NewRecorder()
	request, err := http.NewRequest("GET", url, nil)

	assert.NoError(t, err)
	assert.NotNil(t, request)

	server.Router.ServeHTTP(recorder, request)

	assert.Equal(t, 200, recorder.Result().StatusCode)
}

func TestCreateBook(t *testing.T) {
	ctrl := gomock.NewController(t)

	// Init testServer
	conf := config.InitConfig()
	server := server.NewServer("development")
	server.ServerWrapper(config.InitPostgres(conf))

	defer ctrl.Finish()

	service := mock_service.NewMockBookService(ctrl)
	ctx := context.Background()
	// Stub

	book := &request.CreateBookRequest{
		IDCategory: "5ececdde-26bc-4305-9861-afaaf5291591",
		IDGenre:    "dd5c8932-526a-4e97-9ee7-6f60f3b4848b",
		Title:      "Javier",
		ISBN:       "0091-0213",
		Price:      50000,
	}

	b := &entity.Books{
		ID:         utils.RandomIDGenerator(),
		IDCategory: book.IDCategory,
		IDGenre:    book.IDCategory,
		Title:      book.Title,
		ISBN:       book.ISBN,
		Price:      book.Price,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
	service.EXPECT().CreateBook(ctx, book).Times(1).Return(b, nil).AnyTimes()

	url := "/api/book"
	r := fmt.Sprintf(`{"id_category":"%s", "id_genre":"%s", "title":"%s", "isbn":"%s", "price":%e}`, book.IDCategory, book.IDGenre, book.Title, book.ISBN, book.Price)

	recorder := httptest.NewRecorder()
	request, err := http.NewRequest("POST", url, bytes.NewReader([]byte(r)))
	request.Header.Set("Content-Type", "application/json")
	log.Println("R :", string(r))
	assert.NoError(t, err)
	assert.NotNil(t, request)

	server.Router.ServeHTTP(recorder, request)

	assert.Equal(t, 201, recorder.Result().StatusCode)
}
