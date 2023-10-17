package handler

import (
	"database/sql"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	mocks "github.com/Flo0807/hsfl-master-ai-cloud-engineering/src/auth-service/_mocks"
	"github.com/Flo0807/hsfl-master-ai-cloud-engineering/src/auth-service/pkg/model"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestRegisterHandler(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	userRepository := mocks.NewMockRepository(ctrl)
	hasher := mocks.NewMockHasher(ctrl)
	handler := NewRegisterHandler(userRepository, hasher)

	t.Run("should return 405 METHOD NOT ALLOWED", func(t *testing.T) {
		// given
		w := httptest.NewRecorder()
		r := httptest.NewRequest("GET", "/register", nil)

		// when
		handler.ServeHTTP(w, r)

		// test
		assert.Equal(t, http.StatusMethodNotAllowed, w.Code)
	})

	t.Run("should return 400 BAD REQUEST if payload is not json", func(t *testing.T) {
		tests := []io.Reader{
			nil,
			strings.NewReader(`{"invalid`),
		}

		for _, test := range tests {
			// given
			w := httptest.NewRecorder()
			r := httptest.NewRequest("POST", "/register", test)

			// when
			handler.ServeHTTP(w, r)

			// test
			assert.Equal(t, http.StatusBadRequest, w.Code)
		}
	})

	t.Run("should return 400 BAD REQUEST if payload is not valid", func(t *testing.T) {
		tests := []io.Reader{
			strings.NewReader(`{"email": "email@example.com"}`),
			strings.NewReader(`{"password": "pw"}`),
			strings.NewReader(`{"random": "random"}`),
		}

		for _, test := range tests {
			// given
			w := httptest.NewRecorder()
			r := httptest.NewRequest("POST", "/register", test)

			// when
			handler.ServeHTTP(w, r)

			// test
			assert.Equal(t, http.StatusBadRequest, w.Code)
		}
	})

	t.Run("should return 500 INTERNAL SERVER ERROR if user repository fails", func(t *testing.T) {
		// given
		w := httptest.NewRecorder()
		r := httptest.NewRequest("POST", "/register", strings.NewReader(`{"email": "email@example.com", "password": "password"}`))

		userRepository.
			EXPECT().
			FindUserByEmail("email@example.com").
			Return(nil, errors.New("database error"))

		// when
		handler.ServeHTTP(w, r)

		// test
		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})

	t.Run("should return 409 CONFLICT if user already exists", func(t *testing.T) {
		// given
		w := httptest.NewRecorder()
		r := httptest.NewRequest("POST", "/register", strings.NewReader(`{"email": "email@example.com", "password": "password"}`))

		userRepository.
			EXPECT().
			FindUserByEmail("email@example.com").
			Return(&model.DbUser{}, nil)

		// when
		handler.ServeHTTP(w, r)

		// test
		assert.Equal(t, http.StatusConflict, w.Code)
	})

	t.Run("should return 500 INTERNAL SERVER ERROR if hashing password fails", func(t *testing.T) {
		// given
		w := httptest.NewRecorder()
		r := httptest.NewRequest("POST", "/register", strings.NewReader(`{"email": "email@example.com", "password": "password"}`))

		userRepository.
			EXPECT().
			FindUserByEmail("email@example.com").
			Return(nil, sql.ErrNoRows)

		hasher.
			EXPECT().
			Hash([]byte("password")).
			Return(nil, errors.New("hashing error"))

		// when
		handler.ServeHTTP(w, r)

		// test
		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})

	t.Run("should return 500 INTERNAL SERVER ERROR if creating user fails", func(t *testing.T) {
		// given
		w := httptest.NewRecorder()
		r := httptest.NewRequest("POST", "/register", strings.NewReader(`{"email": "email@example.com", "password": "password"}`))

		userRepository.
			EXPECT().
			FindUserByEmail("email@example.com").
			Return(nil, sql.ErrNoRows)

		hasher.
			EXPECT().
			Hash([]byte("password")).
			Return([]byte("hashedPw"), nil)

		userRepository.
			EXPECT().
			CreateUser(gomock.Any()).
			Return(errors.New("database error"))

		// when
		handler.ServeHTTP(w, r)

		// test
		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})

	t.Run("should return 201 CREATED if user is created", func(t *testing.T) {
		// given
		w := httptest.NewRecorder()
		r := httptest.NewRequest("POST", "/register", strings.NewReader(`{"email": "email@example.com", "password": "password"}`))

		userRepository.
			EXPECT().
			FindUserByEmail("email@example.com").
			Return(nil, sql.ErrNoRows)

		hasher.
			EXPECT().
			Hash([]byte("password")).
			Return([]byte("hashedPw"), nil)

		userRepository.
			EXPECT().
			CreateUser(gomock.Any()).
			Return(nil)

		// when
		handler.ServeHTTP(w, r)

		// test
		assert.Equal(t, http.StatusCreated, w.Code)
	})
}
