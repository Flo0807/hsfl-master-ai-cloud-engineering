package middleware

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Flo0807/hsfl-master-ai-cloud-engineering/lib/rpc/auth"
	mocks "github.com/Flo0807/hsfl-master-ai-cloud-engineering/lib/rpc/auth/_mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"gotest.tools/v3/assert"
)

func TestCreateAuthMiddleware(t *testing.T) {
	mockAuthServiceClient := &mocks.MockAuthServiceClient{}
	middleware := CreateAuthMiddleware(mockAuthServiceClient)

	req, err := http.NewRequest("GET", "/", nil)
	require.NoError(t, err)

	handler := middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))

	t.Run("should return 401 UNAUTHORIZED if no token is provided", func(t *testing.T) {
		// given
		recorder := httptest.NewRecorder()

		// when
		handler.ServeHTTP(recorder, req)

		// test
		assert.Equal(t, http.StatusUnauthorized, recorder.Code)
	})

	t.Run("should return 401 UNAUTHORIZED if token is invalid", func(t *testing.T) {
		// given
		recorder := httptest.NewRecorder()

		// when
		req.Header.Set("Authorization", "invalid")
		handler.ServeHTTP(recorder, req)

		// test
		assert.Equal(t, http.StatusUnauthorized, recorder.Code)
	})

	t.Run("should return 200 OK if token is valid", func(t *testing.T) {
		// given
		recorder := httptest.NewRecorder()

		// when
		validToken := "Bearer valid_token"
		req.Header.Set("Authorization", validToken)

		mockAuthServiceClient.
			On("ValidateToken", mock.Anything, &auth.ValidateTokenRequest{
				Token: strings.Split(validToken, " ")[1],
			}, mock.Anything).
			Return(&auth.ValidateTokenResponse{
				Valid: true,
				User: &auth.User{
					Id:       1,
					Username: "test",
				},
			}, nil)

		handler.ServeHTTP(recorder, req)

		// test
		assert.Equal(t, http.StatusOK, recorder.Code)
	})
}
