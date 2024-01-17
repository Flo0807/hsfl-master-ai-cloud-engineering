package feed

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Flo0807/hsfl-master-ai-cloud-engineering/lib/containerhelpers"
	"github.com/Flo0807/hsfl-master-ai-cloud-engineering/lib/rpc/bulletin-board/rpc/bulletin_board"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type MockBulletinBoardServiceClient struct{}

func (m *MockBulletinBoardServiceClient) GetPosts(ctx context.Context, req *bulletin_board.Request, opts ...grpc.CallOption) (*bulletin_board.Feed, error) {
	post1CreatedAt := timestamppb.New(time.Now())
	post2CreatedAt := timestamppb.New(time.Now().Add(-1 * time.Hour))

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

	assert.Equal(t, "Post 1 Title", response.Posts[0].Title)
	assert.Equal(t, "Post 1 Content", response.Posts[0].Content)
}

func TestDefaultController_GetFeed_ErrorFromClient(t *testing.T) {
	mockClient := &MockBulletinBoardServiceClient{}
	ctrl := NewDefaultController(mockClient)

	req, err := http.NewRequest("GET", "/feed?amount=-1", nil)
	assert.NoError(t, err)

	w := httptest.NewRecorder()

	ctrl.GetFeed(w, req)
	assert.Equal(t, http.StatusInternalServerError, w.Code)

	var errorResponse map[string]string
	err = json.Unmarshal(w.Body.Bytes(), &errorResponse)

	assert.NoError(t, err)
	assert.Contains(t, errorResponse, "error")

	expectedErrorMessage := "Internal Server Error"
	assert.Equal(t, expectedErrorMessage, errorResponse["error"])
}

func TestDefaultController_GetHealth(t *testing.T) {
	// Setup
	ctrl := NewDefaultController(nil)

	req, err := http.NewRequest("GET", "/health", nil)
	assert.NoError(t, err)

	w := httptest.NewRecorder()

	// Test
	ctrl.GetHealth(w, req)

	// Assertions
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Empty(t, w.Body.String(), "Expected an empty response body for health check")
}

func TestIntegration(t *testing.T) {

	//Setup and start containers to check if grpc connection can be established

	postgres, err := containerhelpers.StartPostgres()
	if err != nil {
		t.Fatalf("could not start postgres container: %s", err.Error())
	}
	fmt.Println(postgres.GetContainerID())

	posgtresIp, err := postgres.ContainerIP(context.Background())

	if err != nil {
		t.Fatalf("could not get postgres container ip: %s", err.Error())
	}

	authService, err := containerhelpers.StartAuthService(posgtresIp)
	if err != nil {
		t.Fatalf("could not start auth container: %s", err.Error())
	}
	fmt.Println(authService.GetContainerID())

	bulletinContainer, err := containerhelpers.StartBulletinService(posgtresIp)
	if err != nil {
		t.Fatalf("could not start bulletin-board container: %s", err.Error())
	}
	fmt.Println(bulletinContainer.GetContainerID())

	bulletinBoardIp, err := bulletinContainer.ContainerIP(context.Background())

	if err != nil {
		t.Fatalf("could not get bulletin-board container ip: %s", err.Error())
	}

	feedContainer, err := containerhelpers.StartFeedService(bulletinBoardIp)
	if err != nil {
		t.Fatalf("could not start postgres container: %s", err.Error())
	}
	fmt.Println(feedContainer.GetContainerID())

	feedIp, err := feedContainer.ContainerIP(context.Background())

	if err != nil {
		t.Fatalf("could not get feed container ip: %s", err.Error())
	}

	feedUrl := fmt.Sprintf("http://%s:%s/feed/feed?amount=1", feedIp, "3000")

	resp, err := http.Get(feedUrl)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.NoError(t, err)

}
