package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/Flo0807/hsfl-master-ai-cloud-engineering/bulletin-board-service/api/handler"
	"github.com/Flo0807/hsfl-master-ai-cloud-engineering/bulletin-board-service/api/router"
	"github.com/Flo0807/hsfl-master-ai-cloud-engineering/bulletin-board-service/api/rpc"
	"github.com/Flo0807/hsfl-master-ai-cloud-engineering/bulletin-board-service/config"
	"github.com/Flo0807/hsfl-master-ai-cloud-engineering/bulletin-board-service/models"
	"github.com/Flo0807/hsfl-master-ai-cloud-engineering/bulletin-board-service/repository"
	"github.com/Flo0807/hsfl-master-ai-cloud-engineering/bulletin-board-service/service"
	"github.com/Flo0807/hsfl-master-ai-cloud-engineering/lib/rpc/auth"
	proto "github.com/Flo0807/hsfl-master-ai-cloud-engineering/lib/rpc/bulletin-board/rpc/bulletin_board"
	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	godotenv.Load()

	cfg := config.Config{}

	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("error while parsing enviroment variables: %s", err.Error())
	}

	db, err := gorm.Open(postgres.Open(cfg.Database.Dsn()), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	err = db.AutoMigrate(&models.Post{})
	if err != nil {
		log.Fatal("Failed to migrate database: ", err)
	}

	postRepo := repository.NewPostPsqlRepository(db)
	postService := service.NewPostService(postRepo)
	postHandler := handler.NewPostHandler(postService)

	healthHandler := handler.NewHealthHandler()

	grpcConn, err := grpc.Dial(cfg.AuthServiceUrlGrpc, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("error while connecting to auth service: %s", err.Error())
	}

	authServiceClient := auth.NewAuthServiceClient(grpcConn)

	r := router.NewRouter(healthHandler, postHandler, authServiceClient)

	go func() {
		log.Printf("Starting HTTP server on port %s", cfg.HttpServerPort)

		addr := fmt.Sprintf("0.0.0.0:%s", cfg.HttpServerPort)
		if err := http.ListenAndServe(addr, r); err != nil {
			log.Fatalf("error while listen and serve: %s", err.Error())
		}
	}()
	//go func() {
	log.Printf("Starting gRPC server on port %s", cfg.GrpcServerPort)

	listener, err := net.Listen("tcp", ":"+cfg.GrpcServerPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	bulletinBoardServiceServer := rpc.NewBulletinBoardServiceServer(postHandler)
	proto.RegisterBulletinBoardServiceServer(grpcServer, bulletinBoardServiceServer)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err.Error())
	}
	//}()
}
