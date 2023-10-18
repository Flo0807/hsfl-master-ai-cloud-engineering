package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/Flo0807/hsfl-master-ai-cloud-engineering/src/auth-service/crypto"
	"github.com/Flo0807/hsfl-master-ai-cloud-engineering/src/auth-service/pkg/model"
	"github.com/Flo0807/hsfl-master-ai-cloud-engineering/src/auth-service/user"
)

type registerRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterHandler struct {
	userRepository user.Repository
	hasher         crypto.Hasher
}

func (r *registerRequest) isValid() bool {
	return r.Email != "" && r.Password != ""
}

func NewRegisterHandler(
	userRepository user.Repository,
	hasher crypto.Hasher,
) *RegisterHandler {
	return &RegisterHandler{userRepository, hasher}
}

func (h *RegisterHandler) Register(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var request registerRequest

		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if !request.isValid() {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		_, err := h.userRepository.FindUserByEmail(request.Email)

		if err != nil && err != sql.ErrNoRows {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if err != sql.ErrNoRows {
			w.WriteHeader(http.StatusConflict)
			return
		}

		hashedPw, err := h.hasher.Hash([]byte(request.Password))

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		user := &model.DbUser{
			Email:    request.Email,
			Password: hashedPw,
		}

		err = h.userRepository.CreateUser(user)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
