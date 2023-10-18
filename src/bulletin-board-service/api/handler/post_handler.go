package handler

import (
	"encoding/json"
	"github.com/Flo0807/hsfl-master-ai-cloud-engineering/bulletin-board-service/service"
	"io"
	"net/http"
	"strconv"

	"github.com/Flo0807/hsfl-master-ai-cloud-engineering/bulletin-board-service/models"
	"github.com/gorilla/mux"
)

// PostHandler handles HTTP requests for Post
type PostHandler struct {
	PostService *service.PostService
}

// NewPostHandler creates a new PostHandler
func NewPostHandler(service *service.PostService) *PostHandler {
	return &PostHandler{PostService: service}
}

// CreatePost handles the creation of a new post
func (h *PostHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
	var post models.Post
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&post); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}
	}(r.Body)

	h.PostService.Create(&post)
	respondWithJSON(w, http.StatusCreated, post)
}

// GetPosts handles the retrieval of all posts
func (h *PostHandler) GetPosts(w http.ResponseWriter, _ *http.Request) {
	posts := h.PostService.GetAll()
	respondWithJSON(w, http.StatusOK, posts)
}

// GetPost handles the retrieval of a post by ID
func (h *PostHandler) GetPost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := convertToUint(params["id"])
	post := h.PostService.GetByID(id)
	respondWithJSON(w, http.StatusOK, post)
}

// UpdatePost handles the update of a post by ID
func (h *PostHandler) UpdatePost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := convertToUint(params["id"])

	var post models.Post
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&post); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}
	}(r.Body)

	existingPost := h.PostService.GetByID(id)
	if existingPost.ID == 0 {
		respondWithError(w, http.StatusNotFound, "Post not found")
		return
	}

	post.ID = id
	h.PostService.Update(&post)
	respondWithJSON(w, http.StatusOK, post)
}

// DeletePost handles the deletion of a post by ID
func (h *PostHandler) DeletePost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := convertToUint(params["id"])

	post := h.PostService.GetByID(id)
	if post.ID == 0 {
		respondWithError(w, http.StatusNotFound, "Post not found")
		return
	}

	h.PostService.Delete(&post)
	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

// Helper function to respond with JSON
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, err = w.Write(response)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
}

// Helper function to respond with an error message
func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

// Helper function to convert string to uint
func convertToUint(s string) uint {
	// Add proper error handling based on your requirements
	result, _ := strconv.ParseUint(s, 10, 64)
	return uint(result)
}
