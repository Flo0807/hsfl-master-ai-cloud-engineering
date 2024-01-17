package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/Flo0807/hsfl-master-ai-cloud-engineering/src/auth-service/auth"
	"github.com/Flo0807/hsfl-master-ai-cloud-engineering/src/auth-service/crypto"
	"github.com/Flo0807/hsfl-master-ai-cloud-engineering/src/auth-service/user"
)

type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type loginResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

type LoginHandler struct {
	userRepository    user.Repository
	hasher            crypto.Hasher
	jwtTokenGenerator auth.TokenGenerator
}

func NewLoginHandler(
	userRepository user.Repository,
	hasher crypto.Hasher,
	jwtTokenGenerator auth.TokenGenerator,
) *LoginHandler {
	return &LoginHandler{userRepository, hasher, jwtTokenGenerator}
}

func (h *LoginHandler) Login(w http.ResponseWriter, r *http.Request) {
	var request loginRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !request.isValid() {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := h.userRepository.FindUserByName(request.Username)

	if err != nil {
		if err == sql.ErrNoRows {
			w.Header().Add("WWW-Authenticate", "Basic realm=Restricted")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		w.Header().Add("WWW-Authenticate", "Basic realm=Restricted")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	valid := h.hasher.Validate([]byte(request.Password), user.Password)

	if !valid {
		w.Header().Add("WWW-Authenticate", "Basic realm=Restricted")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	expiration := time.Duration(3600) * time.Second

	claims := map[string]interface{}{
		"sub":      user.ID,
		"username": user.Username,
		"exp":      time.Now().Add(expiration).Unix(),
	}

	token, err := h.jwtTokenGenerator.GenerateToken(claims)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(loginResponse{
		AccessToken: token,
		TokenType:   "Bearer",
		ExpiresIn:   int(expiration.Seconds()),
	})
}

func (r *loginRequest) isValid() bool {
	return r.Username != "" && r.Password != ""
}
