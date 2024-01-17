package models

import (
	"time"
)

type Post struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `sql:"index"`

	Title   string `json:"title"`
	Content string `json:"content"`
}
