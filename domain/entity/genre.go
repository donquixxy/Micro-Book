package entity

import "time"

type Genre struct {
	ID        string    `json:"id" gorm:"column:id"`
	Genre     string    `json:"genre" gorm:"column:genre"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`
}

func (g *Genre) String() string {
	return "genre"
}
