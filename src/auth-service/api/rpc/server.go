package rpc

import (
	"context"
	"log"

	proto "github.com/Flo0807/hsfl-master-ai-cloud-engineering/lib/rpc/auth"
	"github.com/Flo0807/hsfl-master-ai-cloud-engineering/src/auth-service/auth"
	"github.com/Flo0807/hsfl-master-ai-cloud-engineering/src/auth-service/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthServiceServer struct {
	proto.UnimplementedAuthServiceServer
	userRepository user.Repository
	tokenGenerator auth.TokenGenerator
}

func NewAuthServiceServer(userRepository user.Repository, tokenGenerator auth.TokenGenerator) *AuthServiceServer {
	return &AuthServiceServer{
		userRepository: userRepository,
		tokenGenerator: tokenGenerator,
	}
}

func (a *AuthServiceServer) ValidateToken(_ context.Context, request *proto.ValidateTokenRequest) (*proto.ValidateTokenResponse, error) {
	claims, err := a.tokenGenerator.ValidateToken(request.Token)

	if err != nil {
		log.Println("Verification failed: ", err.Error())
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	username, exists := claims["username"].(string)

	if !exists {
		log.Println("Verification failed: ", err.Error())
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	user, err := a.userRepository.FindUserByName(username)

	if err != nil {
		log.Println("Verification failed: ", err.Error())
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	return &proto.ValidateTokenResponse{
		Valid: true,
		User: &proto.User{
			Id:       uint64(user.ID),
			Username: user.Username,
		},
	}, nil
}
