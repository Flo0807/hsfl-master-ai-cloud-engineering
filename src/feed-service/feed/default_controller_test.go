package feed

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Flo0807/hsfl-master-ai-cloud-engineering/lib/rpc/bulletin-board/rpc/bulletin_board"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// MockBulletinBoardServiceClient is a mock implementation of BulletinBoardServiceClient
type MockBulletinBoardServiceClient struct{}

func (m *MockBulletinBoardServiceClient) GetPosts(ctx context.Context, req *bulletin_board.Request, opts ...grpc.CallOption) (*bulletin_board.Feed, error) {
	// Mocking response for testing purposes
	post1CreatedAt := timestamppb.New(time.Now())
	post2CreatedAt := timestamppb.New(time.Now().Add(-1 * time.Hour)) // one hour ago

	return &bulletin_board.Feed{
		Posts: []*bulletin_board.Post{
			{
				ID:        1,
				CreatedAt: post1CreatedAt,
				UpdatedAt: post1CreatedAt,
				DeletedAt: nil,
				Title:     "Post 1 Title",
				Content:   "Post 1 Content",
			},
			{
				ID:        2,
				CreatedAt: post2CreatedAt,
				UpdatedAt: post2CreatedAt,
				DeletedAt: nil,
				Title:     "Post 2 Title",
				Content:   "Post 2 Content",
			},
		},
	}, nil
}

func TestDefaultController_GetFeed(t *testing.T) {
	// Setup
	mockClient := &MockBulletinBoardServiceClient{}
	ctrl := NewDefaultController(mockClient)

	req, err := http.NewRequest("GET", "/feed?amount=2", nil)
	assert.NoError(t, err)

	w := httptest.NewRecorder()

	// Test
	ctrl.GetFeed(w, req)

	// Assertions
	assert.Equal(t, http.StatusOK, w.Code)

	var response bulletin_board.Feed
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Len(t, response.Posts, 2)

	// Add more assertions based on the structure of your Post message
	assert.Equal(t, "Post 1 Title", response.Posts[0].Title)
	assert.Equal(t, "Post 1 Content", response.Posts[0].Content)
	// Add similar assertions for other fields as needed
}

func TestDefaultController_GetFeed_ErrorFromClient(t *testing.T) {
	// Setup
	mockClient := &MockBulletinBoardServiceClient{}
	ctrl := NewDefaultController(mockClient)

	// Intentionally set the amount to a negative value to simulate an error from the client
	req, err := http.NewRequest("GET", "/feed?amount=-1", nil)
	assert.NoError(t, err)

	w := httptest.NewRecorder()

	// Test
	ctrl.GetFeed(w, req)

	// Assertions
	assert.Equal(t, http.StatusInternalServerError, w.Code)

	var errorResponse map[string]string
	err = json.Unmarshal(w.Body.Bytes(), &errorResponse)
	assert.NoError(t, err)

	// Check if the "error" field exists in the response
	assert.Contains(t, errorResponse, "error")

	// Check the error message
	expectedErrorMessage := "Internal Server Error"
	assert.Equal(t, expectedErrorMessage, errorResponse["error"])
}

func TestDefaultController_GetHealth(t *testing.T) {
	// Setup
	ctrl := NewDefaultController(nil) // Pass a nil client for simplicity in health check

	req, err := http.NewRequest("GET", "/health", nil)
	assert.NoError(t, err)

	w := httptest.NewRecorder()

	// Test
	ctrl.GetHealth(w, req)

	// Assertions
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Empty(t, w.Body.String(), "Expected an empty response body for health check")
}
