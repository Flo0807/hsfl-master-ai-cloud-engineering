package handler

import (
	"database/sql"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"go.uber.org/mock/gomock"

	mocks "github.com/Flo0807/hsfl-master-ai-cloud-engineering/src/auth-service/_mocks"
	"github.com/Flo0807/hsfl-master-ai-cloud-engineering/src/auth-service/pkg/model"
	"github.com/stretchr/testify/assert"
)

func TestLoginHandler(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	userRepository := mocks.NewMockRepository(ctrl)
	hasher := mocks.NewMockHasher(ctrl)
	jwtTokenGenerator := mocks.NewMockTokenGenerator(ctrl)
	handler := NewLoginHandler(userRepository, hasher, jwtTokenGenerator)

	t.Run("should return 400 BAD REQUEST if payload is not json", func(t *testing.T) {
		tests := []io.Reader{
			nil,
			strings.NewReader(`{"invalid`),
		}

		for _, test := range tests {
			// given
			w := httptest.NewRecorder()
			r := httptest.NewRequest("POST", "/login", test)

			// when
			handler.Login(w, r)

			// test
			assert.Equal(t, http.StatusBadRequest, w.Code)
		}
	})

	t.Run("should return 400 BAD REQUEST if payload is not valid", func(t *testing.T) {
		tests := []io.Reader{
			strings.NewReader(`{"username": "user"}`),
			strings.NewReader(`{"password": "pw"}`),
			strings.NewReader(`{"random": "random"}`),
		}

		for _, test := range tests {
			// given
			w := httptest.NewRecorder()
			r := httptest.NewRequest("POST", "/login", test)

			// when
			handler.Login(w, r)

			// test
			assert.Equal(t, http.StatusBadRequest, w.Code)
		}
	})

	t.Run("should return 401 UNAUTHORIZED if user does not exist", func(t *testing.T) {
		// given
		w := httptest.NewRecorder()
		r := httptest.NewRequest("POST", "/login", strings.NewReader(`{"username": "user", "password": "password"}`))

		userRepository.
			EXPECT().
			FindUserByName("user").
			Return(nil, sql.ErrNoRows)

		// when
		handler.Login(w, r)

		// test
		assert.Equal(t, "Basic realm=Restricted", w.Header().Get("WWW-Authenticate"))
		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})

	t.Run("should return 401 UNAUTHORIZED if password is invalid", func(t *testing.T) {
		// given
		w := httptest.NewRecorder()
		r := httptest.NewRequest("POST", "/login", strings.NewReader(`{"username": "user", "password": "password"}`))

		userRepository.
			EXPECT().
			FindUserByName("user").
			Return(&model.DbUser{Password: []byte("random")}, nil)

		hasher.
			EXPECT().
			Validate([]byte("password"), []byte("random")).
			Return(false)

		// when
		handler.Login(w, r)

		// test
		assert.Equal(t, "Basic realm=Restricted", w.Header().Get("WWW-Authenticate"))
		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})

	t.Run("should return 500 INTERNAL SERVER ERROR if user repository fails", func(t *testing.T) {
		// given
		w := httptest.NewRecorder()
		r := httptest.NewRequest("POST", "/login", strings.NewReader(`{"username": "user", "password": "password"}`))

		userRepository.
			EXPECT().
			FindUserByName("user").
			Return(nil, errors.New("database error"))

		// when
		handler.Login(w, r)

		// test
		assert.Equal(t, "Basic realm=Restricted", w.Header().Get("WWW-Authenticate"))
		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})

	t.Run("should return 200 OK", func(t *testing.T) {
		// given
		w := httptest.NewRecorder()
		r := httptest.NewRequest("POST", "/login", strings.NewReader(`{"username": "user", "password": "password"}`))

		userRepository.
			EXPECT().
			FindUserByName("user").
			Return(&model.DbUser{ID: 0, Username: "user", Password: []byte("password")}, nil)

		hasher.
			EXPECT().
			Validate([]byte("password"), []byte("password")).
			Return(true)

		jwtTokenGenerator.
			EXPECT().
			GenerateToken(gomock.Any()).
			Return("token", nil)

		// when
		handler.Login(w, r)

		// test
		res := w.Result()
		var response map[string]interface{}
		err := json.NewDecoder(res.Body).Decode(&response)

		assert.NoError(t, err)

		assert.Equal(t, "token", response["access_token"])
		assert.Equal(t, "Bearer", response["token_type"])
		assert.Equal(t, http.StatusOK, w.Code)
	})
}
