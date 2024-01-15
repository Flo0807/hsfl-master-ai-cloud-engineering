package repository

import "github.com/Flo0807/hsfl-master-ai-cloud-engineering/bulletin-board-service/models"

type PostRepository interface {
	Create(post *models.Post)
	FindAll(take int64, skip int64) PostPage
	FindByID(id uint) models.Post
	Count() int64
	Update(post *models.Post)
	Delete(post *models.Post)
}
