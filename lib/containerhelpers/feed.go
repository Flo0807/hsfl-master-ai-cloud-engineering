package containerhelpers

import (
	"context"
	"fmt"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func StartFeedService(bulletinBoardHost string) (testcontainers.Container, error) {
	req := testcontainers.ContainerRequest{
		Image:        "ghcr.io/flo0807/hsfl-master-ai-cloud-engineering/feed-service:main",
		ExposedPorts: []string{"3000"},
		Env: map[string]string{
			"HTTP_SERVER_PORT":                "3000",
			"BULLETIN_BOARD_SERVICE_URL_GRPC": fmt.Sprintf("%s:50052", bulletinBoardHost),
		},
		WaitingFor: wait.ForListeningPort("3000"),
	}

	return testcontainers.GenericContainer(context.Background(), testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
}
