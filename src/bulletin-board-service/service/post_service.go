package service

import (
	"github.com/Flo0807/hsfl-master-ai-cloud-engineering/bulletin-board-service/models"
	"github.com/Flo0807/hsfl-master-ai-cloud-engineering/bulletin-board-service/repository"
)

type PostService interface {
	Create(post *models.Post)
	GetAll(take int64, skip int64) repository.PostPage
	GetByID(id uint) models.Post
	Count() int64
	Update(post *models.Post)
	Delete(post *models.Post)
}
