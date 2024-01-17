package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/Flo0807/hsfl-master-ai-cloud-engineering/bulletin-board-service/repository"
	"github.com/Flo0807/hsfl-master-ai-cloud-engineering/bulletin-board-service/service"
	"golang.org/x/sync/singleflight"

	"github.com/Flo0807/hsfl-master-ai-cloud-engineering/bulletin-board-service/models"
)

// PostHandler handles HTTP requests for Post
type PostHandler struct {
	PostService service.PostService
	g           *singleflight.Group
}

// NewPostHandler creates a new PostHandler
func NewPostHandler(service service.PostService) *PostHandler {
	g := &singleflight.Group{}
	return &PostHandler{service, g}
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
func (h *PostHandler) GetPosts(w http.ResponseWriter, r *http.Request) {
	take := int64(10) // Standardwert für Anzahl der Datensätze pro Seite als uint
	page := int64(1)  // Standardwert für die aktuelle Seite als uint

	takeParam := r.FormValue("take")
	pageParam := r.FormValue("page")

	// Überprüfen und Parsen der optionalen Parameter
	if takeParam != "" {
		takeValue, err := strconv.ParseInt(takeParam, 10, 0)
		if err == nil {
			take = takeValue
		}
	}

	if pageParam != "" {
		pageValue, err := strconv.ParseInt(pageParam, 10, 0)
		if err == nil {
			page = pageValue
		}
	}

	// Berechnung des Datensatz-Offsets basierend auf der aktuellen Seite
	skip := (page - 1) * take

	// Verwende `take` und `skip` in deinem Service, um die Daten zu paginieren
	postPage := h.PostService.GetAll(take, skip)
	respondWithJSON(w, http.StatusOK, postPage)
}

func (h *PostHandler) GetPostsRequestCoalescing(w http.ResponseWriter, r *http.Request) {
	take := int64(10) // default value for number of records per page as uint
	page := int64(1)  // default value for current page as uint

	takeParam := r.FormValue("take")
	pageParam := r.FormValue("page")

	if takeParam != "" {
		takeValue, err := strconv.ParseInt(takeParam, 10, 0)
		if err == nil {
			take = takeValue
		}
	}

	if pageParam != "" {
		pageValue, err := strconv.ParseInt(pageParam, 10, 0)
		if err == nil {
			page = pageValue
		}
	}

	skip := (page - 1) * take

	msg, err, _ := h.g.Do(fmt.Sprint(take, skip), func() (interface{}, error) {
		postPage := h.PostService.GetAll(take, skip)

		return postPage, nil
	})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusOK, msg.(repository.PostPage))
}

// GetPost handles the retrieval of a post by ID
func (h *PostHandler) GetPost(w http.ResponseWriter, r *http.Request) {
	idString := r.Context().Value("id").(string)
	id := convertToUint(idString)

	msg, err, _ := h.g.Do(idString, func() (interface{}, error) {
		post := h.PostService.GetByID(id)

		return post, nil
	})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	post := msg.(models.Post)

	if post.ID == 0 {
		respondWithError(w, http.StatusNotFound, "Post not found")
		return
	}
	respondWithJSON(w, http.StatusOK, post)
}

// UpdatePost handles the update of a post by ID
func (h *PostHandler) UpdatePost(w http.ResponseWriter, r *http.Request) {
	idString := r.Context().Value("id").(string)
	id := convertToUint(idString)

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
	idString := r.Context().Value("id").(string)
	id := convertToUint(idString)

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
