package rpc

import (
	"context"

	"github.com/Flo0807/hsfl-master-ai-cloud-engineering/bulletin-board-service/api/handler"
	"github.com/Flo0807/hsfl-master-ai-cloud-engineering/bulletin-board-service/models"
	"github.com/Flo0807/hsfl-master-ai-cloud-engineering/lib/rpc/bulletin-board/rpc/bulletin_board"
	proto "github.com/Flo0807/hsfl-master-ai-cloud-engineering/lib/rpc/bulletin-board/rpc/bulletin_board"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type BulletinBoardServiceServer struct {
	proto.UnimplementedBulletinBoardServiceServer
	postHandler handler.PostHandler
}

func NewBulletinBoardServiceServer(postHandler *handler.PostHandler) *BulletinBoardServiceServer {
	return &BulletinBoardServiceServer{
		postHandler: *postHandler,
	}
}

func (a *BulletinBoardServiceServer) GetPosts(context.Context, *bulletin_board.Request) (*proto.Feed, error) {
	claims := a.postHandler.PostService.GetAll(int64(10), int64(1))

	var newArray []*proto.Post

	for _, claims := range claims.Records {
		// Transform each old object into a new object and append to the new array
		newObj := transformFunction(claims)
		newArray = append(newArray, newObj)
	}
	return &proto.Feed{
		Posts: newArray,
	}, nil
}

func transformFunction(oldObj models.Post) *proto.Post {
	// Perform the transformation logic here
	// For example, you can create a new NewType object based on the fields of OldType
	newObj := proto.Post{
		ID:        uint64(oldObj.ID),
		CreatedAt: timestamppb.New(oldObj.CreatedAt),
		UpdatedAt: timestamppb.New(oldObj.UpdatedAt),
		Title:     oldObj.Title,
		Content:   oldObj.Content,
	}

	return &newObj
}
