package entity

import "time"

type Books struct {
	ID         string    `json:"id" gorm:"column:id;primaryKey"`
	IDCategory string    `json:"id_category" gorm:"column:id_category;index"`
	IDGenre    string    `json:"id_genre" gorm:"column:id_genre;index"`
	Title      string    `json:"title" gorm:"column:title"`
	ISBN       string    `json:"isbn" gorm:"column:isbn"`
	Price      float64   `json:"price" gorm:"column:price"`
	CreatedAt  time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"column:updated_at"`
}

func (*Books) TableName() string {
	return "book"
}
