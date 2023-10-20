package feed

import (
	"encoding/json"
	"net/http"
)

type DefaultController struct {
}

func NewDefaultController() *DefaultController {
	return &DefaultController{}
}

func (ctrl *DefaultController) GetFeed(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://localhost:3000/posts")
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to fetch data")
		return
	}
	respondWithJSON(w, http.StatusOK, resp)

}

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
