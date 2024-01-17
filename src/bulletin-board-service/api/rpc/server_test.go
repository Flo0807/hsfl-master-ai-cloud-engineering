package rpc

import (
	"context"
	"testing"
	"time"

	"github.com/Flo0807/hsfl-master-ai-cloud-engineering/bulletin-board-service/models"
	"github.com/Flo0807/hsfl-master-ai-cloud-engineering/bulletin-board-service/repository"

	"github.com/Flo0807/hsfl-master-ai-cloud-engineering/lib/rpc/bulletin-board/rpc/bulletin_board"
	"github.com/stretchr/testify/assert"
)

type MockPostService struct{}

func (m *MockPostService) GetAll(limit, page int64) repository.PostPage {

	return repository.PostPage{
		Page: repository.Page{
			TotalRecords: 2,
		},
		Records: []models.Post{
			{
				ID:        1,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				Title:     "Test Post 1",
				Content:   "Test Content 1",
			},
			{
				ID:        2,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				Title:     "Test Post 2",
				Content:   "Test Content 2",
			},
		},
	}
}

func TestBulletinBoardServiceServer_GetPosts(t *testing.T) {
	// Setup
	mockPostService := &MockPostService{}
	bulletinBoardService := NewBulletinBoardServiceServer(mockPostService)

	ctx := context.Background()
	request := &bulletin_board.Request{Amount: 10}

	// Test
	response, err := bulletinBoardService.GetPosts(ctx, request)

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, response)

	assert.Len(t, response.Posts, 2)

	// Assertions
	assert.Equal(t, uint64(1), response.Posts[0].ID)
	assert.Equal(t, "Test Post 1", response.Posts[0].Title)
	assert.Equal(t, uint64(2), response.Posts[1].ID)
	assert.Equal(t, "Test Post 2", response.Posts[1].Title)
}

// Mocking the Service methods because they have to be implemented
func (m *MockPostService) Create(post *models.Post) {
	return
}
func (m *MockPostService) Delete(post *models.Post) {
	return
}
func (m *MockPostService) Update(post *models.Post) {
	return
}
func (m *MockPostService) GetByID(int uint) models.Post {
	return models.Post{}
}
