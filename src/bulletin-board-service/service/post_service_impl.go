package service

import (
	"github.com/Flo0807/hsfl-master-ai-cloud-engineering/bulletin-board-service/models"
	"github.com/Flo0807/hsfl-master-ai-cloud-engineering/bulletin-board-service/repository"
	"golang.org/x/sync/singleflight"
)

// PostServiceImpl handles business logic for Post
type PostServiceImpl struct {
	PostRepository repository.PostRepository
	g              *singleflight.Group
}

// NewPostService creates a new PostServiceImpl
func NewPostService(repo repository.PostRepository) *PostServiceImpl {
	g := &singleflight.Group{}

	return &PostServiceImpl{PostRepository: repo, g: g}
}

// Create a new post
func (s *PostServiceImpl) Create(post *models.Post) {
	s.PostRepository.Create(post)
}

// GetAll Get all posts
func (s *PostServiceImpl) GetAll(take int64, skip int64) repository.PostPage {
	return s.PostRepository.FindAll(take, skip)
}

// GetByID Get a post by ID
func (s *PostServiceImpl) GetByID(id uint) models.Post {
	return s.PostRepository.FindByID(id)
}

// Count Get the number of posts
func (s *PostServiceImpl) Count() int64 {
	return s.PostRepository.Count()
}

// Update a post
func (s *PostServiceImpl) Update(post *models.Post) {
	s.PostRepository.Update(post)
}

// Delete a post
func (s *PostServiceImpl) Delete(post *models.Post) {
	s.PostRepository.Delete(post)
}
