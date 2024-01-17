package containerhelpers

import (
	"context"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func StartGatewayService() (testcontainers.Container, error) {
	req := testcontainers.ContainerRequest{
		Image:        "hsfl-master-ai-cloud-engineering-gateway",
		ExposedPorts: []string{"3000:3000"},
		Env: map[string]string{
			"HTTP_SERVER_PORT":           "3000",
			"WEB_SERVICE_URL":            "http://web-service:3000",
			"AUTH_SERVICE_URL":           "http://auth-service:3000",
			"BULLETIN_BOARD_SERVICE_URL": "http://bulletin-board-service:3000",
			"FEED_SERVICE_URL":           "http://feed-service:3000",
		},
		WaitingFor: wait.ForListeningPort("3000"),
	}

	return testcontainers.GenericContainer(context.Background(), testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
}
