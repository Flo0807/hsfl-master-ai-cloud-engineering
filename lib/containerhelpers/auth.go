package containerhelpers

import (
	"context"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func StartAuthService(dbHost string) (testcontainers.Container, error) {
	req := testcontainers.ContainerRequest{
		Image:        "hsfl-master-ai-cloud-engineering-auth-service",
		ExposedPorts: []string{"3000", "50051"},
		Env: map[string]string{
			"HTTP_SERVER_PORT": "3000",
			"GRPC_SERVER_PORT": "50051",
			"JWT_PRIVATE_KEY":  "./key",
			"DB_HOST":          dbHost,
			"DB_PORT":          "5432",
			"DB_USER":          "postgres",
			"DB_PASSWORD":      "password",
			"DB_NAME":          "postgres",
		},
		WaitingFor: wait.ForListeningPort("3000"),
	}

	return testcontainers.GenericContainer(context.Background(), testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
}
