package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Flo0807/hsfl-master-ai-cloud-engineering/feed-service/api/router"
	"github.com/Flo0807/hsfl-master-ai-cloud-engineering/feed-service/config"
	"github.com/Flo0807/hsfl-master-ai-cloud-engineering/feed-service/feed"
	"github.com/Flo0807/hsfl-master-ai-cloud-engineering/lib/rpc/bulletin-board/rpc/bulletin_board"
	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	godotenv.Load()

	cfg := config.Config{}

	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("error while parsing enviroment variables: %s", err.Error())
	}

	grpcConn, err := grpc.Dial(cfg.BuleltinBoardServiceUrlGrpc, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("error while connecting to auth service: %s", err.Error())
	}

	bulletinBoardServiceClient := bulletin_board.NewBulletinBoardServiceClient(grpcConn)
	feedController := feed.NewDefaultController(bulletinBoardServiceClient)
	handler := router.New(feedController)

	log.Printf("Starting HTTP server on port %s", cfg.HttpServerPort)

	addr := fmt.Sprintf("0.0.0.0:%s", cfg.HttpServerPort)
	if err := http.ListenAndServe(addr, handler); err != nil {
		log.Fatalf("error while listen and serve: %s", err.Error())
	}
}
