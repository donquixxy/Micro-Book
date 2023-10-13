package entity

import "time"

type Category struct {
	ID        string    `json:"id" gorm:"column:id;primaryKey"`
	Category  string    `json:"category" gorm:"column:category"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`
}

func (c *Category) TableName() string {
	return "category"
}
