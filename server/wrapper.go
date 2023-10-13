package server

import (
	"micro-book/module/handler"
	"micro-book/module/repository/postgres"
	"micro-book/module/service"

	"gorm.io/gorm"
)

func (s *Server) ServerWrapper(db *gorm.DB) {

	// Repository
	booksRepository := postgres.NewBooksRepository(db)
	bookService := service.NewBookService(booksRepository)

	bookHandler := handler.NewBooksHandler(bookService)
	s.registerBookRoute(bookHandler)
}

func (s *Server) registerBookRoute(h handler.BooksHandler) {
	group := s.Router.Group("/api")

	group.POST("/book", h.CreateBook)
	group.POST("/category", h.CreateCategory)
	group.POST("/genre", h.CreateGenre)
	group.GET("/book/:id", h.GetByID)
	group.GET("/books", h.GetAll)
	group.GET("/book/category/:id", h.GetByIDCategory)
}
