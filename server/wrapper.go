package server

import (
	"micro-book/module/repository/postgres"

	"gorm.io/gorm"
)

func ServerWrapper(db *gorm.DB) {

	// Repository
	booksRepository := postgres.NewBooksRepository(db)

	_ = booksRepository
}
