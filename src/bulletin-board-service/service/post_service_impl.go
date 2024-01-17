package service

import (
	"github.com/Flo0807/hsfl-master-ai-cloud-engineering/bulletin-board-service/models"
	"github.com/Flo0807/hsfl-master-ai-cloud-engineering/bulletin-board-service/repository"
)

type PostServiceImpl struct {
	PostRepository repository.PostRepository
}

func NewPostService(repo repository.PostRepository) *PostServiceImpl {
	return &PostServiceImpl{PostRepository: repo}
}

func (s *PostServiceImpl) Create(post *models.Post) {
	s.PostRepository.Create(post)
}

func (s *PostServiceImpl) GetAll(take int64, skip int64) repository.PostPage {
	return s.PostRepository.FindAll(take, skip)
}

func (s *PostServiceImpl) GetByID(id uint) models.Post {
	return s.PostRepository.FindByID(id)
}

func (s *PostServiceImpl) Update(post *models.Post) {
	s.PostRepository.Update(post)
}

func (s *PostServiceImpl) Delete(post *models.Post) {
	s.PostRepository.Delete(post)
}
