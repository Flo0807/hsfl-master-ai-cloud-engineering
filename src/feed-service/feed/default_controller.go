package feed

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Flo0807/hsfl-master-ai-cloud-engineering/lib/rpc/bulletin-board/rpc/bulletin_board"
)

type DefaultController struct {
	client bulletin_board.BulletinBoardServiceClient
}

func NewDefaultController(bulletinBoardClient bulletin_board.BulletinBoardServiceClient) *DefaultController {
	return &DefaultController{
		client: bulletinBoardClient,
	}
}

func (ctrl *DefaultController) GetFeed(w http.ResponseWriter, r *http.Request) {

	amount := int64(5) // Standardwert für Anzahl der Posts

	amountParam := r.FormValue("amount")

	// Überprüfen und Parsen des optionalen Post Amounts
	if amountParam != "" {
		amountValue, err := strconv.ParseInt(amountParam, 10, 0)
		if err == nil {
			amount = amountValue
		}
	}

	resp, err := ctrl.client.GetPosts(r.Context(), &bulletin_board.Request{Amount: 10})
	if err != nil || amount <= 0 {
		respondWithError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	posts := resp.Posts
	if !(amount >= int64(len(resp.Posts))) {
		posts = resp.Posts[:amount]

	}
	feed := bulletin_board.Feed{
		Posts: posts,
	}
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	respondWithJSON(w, http.StatusOK, feed)

}
func (ctrl *DefaultController) GetHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
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
