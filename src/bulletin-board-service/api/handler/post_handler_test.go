package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	mocks "github.com/Flo0807/hsfl-master-ai-cloud-engineering/bulletin-board-service/_mocks"
	"github.com/Flo0807/hsfl-master-ai-cloud-engineering/bulletin-board-service/models"
	"github.com/Flo0807/hsfl-master-ai-cloud-engineering/bulletin-board-service/repository"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestCreatePost(t *testing.T) {
	// Setup
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockPostService(ctrl)
	handler := NewPostHandler(mockService)

	t.Run("should return 201 CREATED if request body is valid", func(t *testing.T) {
		tests := []io.Reader{
			strings.NewReader(`{"title": "Test Post", "content": "Test Content"}`),
		}

		for _, test := range tests {
			// Setup
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", "/posts", test)

			// Expectations
			// ToDo: Replace with correct model.Post
			mockService.EXPECT().Create(gomock.Any())

			// Test
			handler.CreatePost(w, req)

			// Assertions
			assert.Equal(t, http.StatusCreated, w.Code)
		}
	})

	t.Run("should return 400 BAD REQUEST if request body is invalid", func(t *testing.T) {
		tests := []io.Reader{
			strings.NewReader(""),             // empty body
			strings.NewReader("invalid json"), // invalid json
		}

		for _, test := range tests {
			// Setup
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", "/posts", test)

			// Test
			handler.CreatePost(w, req)

			// Assertions
			assert.Equal(t, http.StatusBadRequest, w.Code)
		}
	})
}

func TestGetPosts(t *testing.T) {
	// Setup
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockPostService(ctrl)
	handler := NewPostHandler(mockService)

	t.Run("should return 200 OK with list of posts", func(t *testing.T) {
		// Setup
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/posts?take=10&page=1", nil)

		// Expectations
		mockService.EXPECT().GetAll(int64(10), int64(0)).Return(repository.PostPage{
			Page: repository.Page{
				CurrentPage:  1,
				PageSize:     10,
				TotalRecords: 0,
				TotalPages:   0,
			},
			Records: []models.Post{},
		})

		// Test
		handler.GetPosts(w, req)

		// Assertions
		assert.Equal(t, http.StatusOK, w.Code)
	})
}

func TestGetPost(t *testing.T) {
	// Setup
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockPostService(ctrl)
	handler := NewPostHandler(mockService)

	t.Run("should return 200 OK with post", func(t *testing.T) {
		// Setup
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/posts/1", nil)
		req = req.WithContext(context.WithValue(req.Context(), "id", "1"))

		// Expectations
		mockService.EXPECT().GetByID(uint(1)).Return(models.Post{ID: 1})

		// Test
		handler.GetPost(w, req)

		// Assertions
		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("should return 404 NOT FOUND if post does not exist", func(t *testing.T) {
		// Setup
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/posts/1", nil)
		req = req.WithContext(context.WithValue(req.Context(), "id", "1"))

		// Expectations
		mockService.EXPECT().GetByID(uint(1)).Return(models.Post{ID: 0})

		// Test
		handler.GetPost(w, req)

		// Assertions
		assert.Equal(t, http.StatusNotFound, w.Code)
	})
}

func TestUpdatePost(t *testing.T) {
	// Setup
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockPostService(ctrl)
	handler := NewPostHandler(mockService)

	t.Run("should return 200 OK with post", func(t *testing.T) {
		w := httptest.NewRecorder()
		post := models.Post{
			ID:      1,
			Title:   "Test Post",
			Content: "Test Content",
		}
		postJSON, _ := json.Marshal(post)
		req, _ := http.NewRequest("PUT", "/posts/1", bytes.NewBuffer(postJSON))
		req = req.WithContext(context.WithValue(req.Context(), "id", "1"))

		// Expectations
		mockService.EXPECT().GetByID(uint(1)).Return(models.Post{
			ID:      1,
			Title:   "Test Post Old",
			Content: "Test Content Old",
		})
		mockService.EXPECT().Update(&post)

		// Test
		handler.UpdatePost(w, req)

		// Assertions
		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("should return 404 NOT FOUND if post does not exist", func(t *testing.T) {
		w := httptest.NewRecorder()
		post := models.Post{
			Title:   "Test Post",
			Content: "Test Content",
		}
		postJSON, _ := json.Marshal(post)
		req, _ := http.NewRequest("PUT", "/posts/1", bytes.NewBuffer(postJSON))
		req = req.WithContext(context.WithValue(req.Context(), "id", "1"))

		// Expectations
		mockService.EXPECT().GetByID(uint(1)).Return(models.Post{
			ID: 0,
		})

		// Test
		handler.UpdatePost(w, req)

		// Assertions
		assert.Equal(t, http.StatusNotFound, w.Code)
	})

	t.Run("should return 400 BAD REQUEST if request body is invalid", func(t *testing.T) {
		tests := []io.Reader{
			strings.NewReader(""),             // empty body
			strings.NewReader("invalid json"), // invalid json
		}

		for _, test := range tests {
			// Setup
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("PUT", "/posts/1", test)
			req = req.WithContext(context.WithValue(req.Context(), "id", "1"))

			// Test
			handler.UpdatePost(w, req)

			// Assertions
			assert.Equal(t, http.StatusBadRequest, w.Code)
		}
	})
}

func TestDeletePost(t *testing.T) {
	// Setup
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockPostService(ctrl)
	handler := NewPostHandler(mockService)

	t.Run("should return 200 OK with post", func(t *testing.T) {
		w := httptest.NewRecorder()
		post := models.Post{
			ID:      1,
			Title:   "Test Post",
			Content: "Test Content",
		}
		req, _ := http.NewRequest("DELETE", "/posts/1", nil)
		req = req.WithContext(context.WithValue(req.Context(), "id", "1"))

		// Expectations
		mockService.EXPECT().GetByID(uint(1)).Return(post)
		mockService.EXPECT().Delete(&post)

		// Test
		handler.DeletePost(w, req)

		// Assertions
		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("should return 404 NOT FOUND if post does not exist", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("DELETE", "/posts/1", nil)
		req = req.WithContext(context.WithValue(req.Context(), "id", "1"))

		// Expectations
		mockService.EXPECT().GetByID(uint(1)).Return(models.Post{
			ID: 0,
		})

		// Test
		handler.DeletePost(w, req)

		// Assertions
		assert.Equal(t, http.StatusNotFound, w.Code)
	})
}

func TestGetRandom(t *testing.T) {
	// Setup
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockPostService(ctrl)
	handler := NewPostHandler(mockService)

	t.Run("should return 200 OK with post", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/posts/random", nil)

		// Expectations
		mockService.EXPECT().Count().Return(int64(1))
		mockService.EXPECT().GetByID(uint(1)).Return(models.Post{ID: 1})

		// Test
		handler.GetRandom(w, req)

		// Assertions
		assert.Equal(t, http.StatusOK, w.Code)
	})
}

func TestGetRandomRequestCoalescing(t *testing.T) {
	// Setup
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockPostService(ctrl)
	handler := NewPostHandler(mockService)

	t.Run("should return 200 OK with post", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/posts/random-request-coalescing", nil)

		// Expectations
		mockService.EXPECT().Count().Return(int64(1))
		mockService.EXPECT().GetByID(uint(1)).Return(models.Post{ID: 1})

		// Test
		handler.GetRandom(w, req)

		// Assertions
		assert.Equal(t, http.StatusOK, w.Code)
	})
}
