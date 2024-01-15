package service

import (
	"testing"

	mocks "github.com/Flo0807/hsfl-master-ai-cloud-engineering/bulletin-board-service/_mocks"
	"github.com/Flo0807/hsfl-master-ai-cloud-engineering/bulletin-board-service/models"
	"github.com/Flo0807/hsfl-master-ai-cloud-engineering/bulletin-board-service/repository"
	"go.uber.org/mock/gomock"
)

func TestCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockPostRepository(ctrl)
	postService := NewPostService(mockRepo)
	post := models.Post{Title: "Test Post", Content: "Test Content"}

	mockRepo.EXPECT().Create(&post)

	postService.Create(&post)
}

func TestGetAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockPostRepository(ctrl)
	postService := NewPostService(mockRepo)

	mockRepo.EXPECT().FindAll(int64(10), int64(1)).Return(repository.PostPage{
		Page: repository.Page{
			CurrentPage:  1,
			PageSize:     10,
			TotalRecords: 0,
			TotalPages:   0,
		},
		Records: []models.Post{},
	})

	postService.GetAll(int64(10), int64(1))
}

func TestGetByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockPostRepository(ctrl)
	postService := NewPostService(mockRepo)

	mockRepo.EXPECT().FindByID(uint(1)).Return(models.Post{})

	postService.GetByID(uint(1))
}

func TestUpdate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockPostRepository(ctrl)
	postService := NewPostService(mockRepo)
	post := models.Post{Title: "Test Post", Content: "Test Content"}

	mockRepo.EXPECT().Update(&post)

	postService.Update(&post)
}

func TestDelete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockPostRepository(ctrl)
	postService := NewPostService(mockRepo)
	post := models.Post{Title: "Test Post", Content: "Test Content"}

	mockRepo.EXPECT().Delete(&post)

	postService.Delete(&post)
}

func TestCount(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockPostRepository(ctrl)
	postService := NewPostService(mockRepo)

	mockRepo.EXPECT().Count().Return(int64(1))

	postService.Count()
}
