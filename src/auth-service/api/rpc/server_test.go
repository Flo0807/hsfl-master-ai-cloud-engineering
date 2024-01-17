package rpc

import (
	"context"
	"errors"
	"testing"

	proto "github.com/Flo0807/hsfl-master-ai-cloud-engineering/lib/rpc/auth"
	mocks "github.com/Flo0807/hsfl-master-ai-cloud-engineering/src/auth-service/_mocks"
	"github.com/Flo0807/hsfl-master-ai-cloud-engineering/src/auth-service/pkg/model"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestValidateToken(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()

	tokenGenerator := mocks.NewMockTokenGenerator(ctrl)
	userRepository := mocks.NewMockRepository(ctrl)
	server := NewAuthServiceServer(userRepository, tokenGenerator)

	t.Run("should return error if token is invalid", func(t *testing.T) {
		tokenGenerator.
			EXPECT().
			ValidateToken(gomock.Any()).
			Return(nil, errors.New("invalid token"))

		_, err := server.ValidateToken(ctx, &proto.ValidateTokenRequest{Token: "invalid token"})
		assert.Error(t, err)
	})

	t.Run("should return error if user does not exist", func(t *testing.T) {
		tokenGenerator.
			EXPECT().
			ValidateToken(gomock.Any()).
			Return(map[string]interface{}{"username": "nonexistent@example.com"}, nil)

		userRepository.
			EXPECT().
			FindUserByName("nonexistent@example.com").
			Return(nil, errors.New("user does not exist"))

		_, err := server.ValidateToken(ctx, &proto.ValidateTokenRequest{Token: "valid token"})

		assert.Error(t, err)
	})

	t.Run("should return user if token is valid", func(t *testing.T) {
		tokenGenerator.
			EXPECT().
			ValidateToken(gomock.Any()).
			Return(map[string]interface{}{"username": "user"}, nil)

		userRepository.
			EXPECT().
			FindUserByName("user").
			Return(&model.DbUser{ID: 1, Username: "user"}, nil)

		response, err := server.ValidateToken(ctx, &proto.ValidateTokenRequest{Token: "valid token"})

		assert.NoError(t, err)
		assert.Equal(t, uint64(1), response.User.Id)
		assert.Equal(t, "user", response.User.Username)
	})
}
