package handler

import (
	"context"
	"errors"
	"micro-book/domain/request"
	"micro-book/module/service"
	"time"

	"github.com/gin-gonic/gin"
)

type BooksHandler interface {
	CreateBook(g *gin.Context)
	CreateGenre(g *gin.Context)
	CreateCategory(g *gin.Context)
	GetByID(g *gin.Context)
	GetAll(g *gin.Context)
	GetByIDCategory(g *gin.Context)
}

type booksHandler struct {
	bookService service.BookService
}

func NewBooksHandler(
	bookService service.BookService,
) BooksHandler {
	return &booksHandler{
		bookService: bookService,
	}
}

func (h *booksHandler) GetByIDCategory(g *gin.Context) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)

	defer cancel()

	id := g.Param("id")

	result, err := h.bookService.GetByIDCategory(ctx, id)

	if err != nil {
		g.JSON(404, errorResponse(err, "not found"))
		return
	}

	g.JSON(200, successResponse(result, 200))
}

func (h *booksHandler) GetAll(g *gin.Context) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)

	defer cancel()

	result, err := h.bookService.GetAll(ctx)

	if err != nil {
		g.JSON(404, errorResponse(err, "not found"))
		return
	}

	g.JSON(200, successResponse(result, 200))
}

func (h *booksHandler) GetByID(g *gin.Context) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)

	defer cancel()

	id := g.Param("id")

	result, err := h.bookService.GetByID(ctx, id)

	if err != nil {
		g.JSON(404, errorResponse(err, "not found"))
		return
	}

	g.JSON(200, successResponse(result, 200))
}

func (h *booksHandler) CreateBook(g *gin.Context) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)

	defer cancel()

	request, err := request.ReadCreateBookRequest(g)

	if err != nil {
		g.JSON(400, errorResponse(err, "failed getting request in body"))
		return
	}

	result, err := h.bookService.CreateBook(ctx, request)

	if err != nil {
		g.JSON(400, errorResponse(err, "failed creating book"))
		return
	}

	g.JSON(201, successResponse(result, 201))
}

func (h *booksHandler) CreateGenre(g *gin.Context) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)

	defer cancel()

	genre := g.PostForm("genre")

	if len(genre) < 3 {
		g.JSON(400, errorResponse(errors.New("invalid genre length"), "invalid length"))
		return
	}

	result, err := h.bookService.CreateGenre(ctx, genre)

	if err != nil {
		g.JSON(400, errorResponse(err, "failed creating genre"))
		return
	}

	g.JSON(201, successResponse(result, 201))
}

func (h *booksHandler) CreateCategory(g *gin.Context) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)

	defer cancel()

	category := g.PostForm("category")

	if len(category) < 3 {
		g.JSON(400, errorResponse(errors.New("invalid category length"), "invalid length"))
		return
	}

	result, err := h.bookService.CreateCategory(ctx, category)

	if err != nil {
		g.JSON(400, errorResponse(err, "failed creating category"))
		return
	}

	g.JSON(201, successResponse(result, 201))
}

func successResponse(data any, statusCode int) gin.H {
	return gin.H{
		"status": statusCode,
		"data":   data,
		"msg":    "Success",
	}
}

func errorResponse(e error, s string) gin.H {
	return gin.H{
		"error":   e.Error(),
		"message": s,
	}
}
