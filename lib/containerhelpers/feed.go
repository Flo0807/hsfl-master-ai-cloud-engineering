package containerhelpers

import (
	"context"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func StartFeedService() (testcontainers.Container, error) {
	req := testcontainers.ContainerRequest{
		Image:        "hsfl-master-ai-cloud-engineering-feed-service",
		ExposedPorts: []string{"3000"},
		Env: map[string]string{
			"HTTP_SERVER_PORT":                "3000",
			"BULLETIN_BOARD_SERVICE_URL_GRPC": "bulletin-board-service:50052",
		},
		WaitingFor: wait.ForListeningPort("3000"),
	}

	return testcontainers.GenericContainer(context.Background(), testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
}
