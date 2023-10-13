package request

import (
	"log"

	"github.com/gin-gonic/gin"
)

type CreateBookRequest struct {
	IDCategory string  `json:"id_category" form:"id_category" binding:"required"`
	IDGenre    string  `json:"id_genre" form:"id_genre" binding:"required"`
	Title      string  `json:"title" form:"title" binding:"required"`
	ISBN       string  `json:"isbn" form:"isbn" binding:"required"`
	Price      float64 `json:"price" form:"price" binding:"required"`
}

func ReadCreateBookRequest(g *gin.Context) (*CreateBookRequest, error) {
	body := new(CreateBookRequest)

	if bindErr := g.Bind(body); bindErr != nil {
		log.Printf("Failed to bind CreateBookRequest: %v", bindErr)
		return nil, bindErr
	}

	return body, nil
}

//

type UpdateBookRequest struct {
	ID         string  `json:"id" form:"id"`
	IDCategory string  `json:"id_category" form:"id_category"`
	IDGenre    string  `json:"id_genre" form:"id_genre"`
	Title      string  `json:"title" form:"title"`
	ISBN       string  `json:"isbn" form:"isbn"`
	Price      float64 `json:"price" form:"price"`
}

func ReadUpdateBookRequest(g *gin.Context) (*UpdateBookRequest, error) {

	body := new(UpdateBookRequest)

	if bindErr := g.Bind(body); bindErr != nil {
		return nil, bindErr
	}

	return body, nil
}
