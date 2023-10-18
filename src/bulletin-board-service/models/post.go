package models

import (
	"github.com/jinzhu/gorm"
)

// Post struct represents a bulletin board post
type Post struct {
	gorm.Model
	Title   string `json:"title"`
	Content string `json:"content"`
}
