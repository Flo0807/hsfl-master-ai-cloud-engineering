package containerhelpers

import (
	"context"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func StartBulletinService(dbHost string) (testcontainers.Container, error) {
	req := testcontainers.ContainerRequest{
		Image:        "hsfl-master-ai-cloud-engineering-bulletin-board-service",
		ExposedPorts: []string{"3000", "50052"},
		Env: map[string]string{
			"HTTP_SERVER_PORT":      "3000",
			"GRPC_SERVER_PORT":      "50052",
			"DB_HOST":               dbHost,
			"DB_PORT":               "5432",
			"DB_USER":               "postgres",
			"DB_PASSWORD":           "password",
			"DB_NAME":               "postgres",
			"AUTH_SERVICE_URL_GRPC": "auth-service:50051",
		},
		WaitingFor: wait.ForListeningPort("3000"),
	}

	return testcontainers.GenericContainer(context.Background(), testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
}
