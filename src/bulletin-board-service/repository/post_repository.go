package repository

import (
	"github.com/Flo0807/hsfl-master-ai-cloud-engineering/bulletin-board-service/models"
	"gorm.io/gorm"
)

// PostRepository handles database operations for Post
type PostRepository struct {
	DB *gorm.DB
}

// NewPostRepository creates a new PostRepository
func NewPostRepository(db *gorm.DB) *PostRepository {
	return &PostRepository{DB: db}
}

// Create a new post
func (r *PostRepository) Create(post *models.Post) {
	r.DB.Create(post)
}

// FindAll posts
func (r *PostRepository) FindAll() []models.Post {
	var posts []models.Post
	r.DB.Find(&posts)
	return posts
}

// FindByID finds a post by ID
func (r *PostRepository) FindByID(id uint) models.Post {
	var post models.Post
	r.DB.First(&post, id)
	return post
}

// Update a post
func (r *PostRepository) Update(post *models.Post) {
	r.DB.Save(post)
}

// Delete a post
func (r *PostRepository) Delete(post *models.Post) {
	r.DB.Delete(post)
}
