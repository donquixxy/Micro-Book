package entity

import "time"

type Books struct {
	ID         string    `json:"id,omitempty" gorm:"column:id;primaryKey"`
	IDCategory string    `json:"id_category,omitempty" gorm:"column:id_category;index"`
	IDGenre    string    `json:"id_genre,omitempty" gorm:"column:id_genre;index"`
	Title      string    `json:"title,omitempty" gorm:"column:title"`
	ISBN       string    `json:"isbn,omitempty" gorm:"column:isbn"`
	Price      float64   `json:"price,omitempty" gorm:"column:price"`
	CreatedAt  time.Time `json:"created_at,omitempty" gorm:"column:created_at"`
	UpdatedAt  time.Time `json:"updated_at,omitempty" gorm:"column:updated_at"`
	Category   *Category `json:"category,omitempty" gorm:"foreignKey:IDCategory;references:ID"`
	Genre      *Genre    `json:"genre,omitempty" gorm:"foreignKey:IDGenre;references:ID"`
}

func (*Books) TableName() string {
	return "book"
}
