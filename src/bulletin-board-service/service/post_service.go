package service

import (
	"github.com/Flo0807/hsfl-master-ai-cloud-engineering/bulletin-board-service/models"
	"github.com/Flo0807/hsfl-master-ai-cloud-engineering/bulletin-board-service/repository"
)

// PostService handles business logic for Post
type PostService struct {
	PostRepository *repository.PostRepository
}

// NewPostService creates a new PostService
func NewPostService(repo *repository.PostRepository) *PostService {
	return &PostService{PostRepository: repo}
}

// Create a new post
func (s *PostService) Create(post *models.Post) {
	s.PostRepository.Create(post)
}

// GetAll Get all posts
func (s *PostService) GetAll() []models.Post {
	return s.PostRepository.FindAll()
}

// GetByID Get a post by ID
func (s *PostService) GetByID(id uint) models.Post {
	return s.PostRepository.FindByID(id)
}

// Update a post
func (s *PostService) Update(post *models.Post) {
	s.PostRepository.Update(post)
}

// Delete a post
func (s *PostService) Delete(post *models.Post) {
	s.PostRepository.Delete(post)
}
