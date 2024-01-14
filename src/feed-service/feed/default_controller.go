package feed

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type DefaultController struct {
}

func NewDefaultController() *DefaultController {
	return &DefaultController{}
}

func (ctrl *DefaultController) GetFeed(w http.ResponseWriter, r *http.Request) {
	return
}
func (ctrl *DefaultController) GetHealth(w http.ResponseWriter, r *http.Request) {
	fmt.Println("feed")
	return
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
