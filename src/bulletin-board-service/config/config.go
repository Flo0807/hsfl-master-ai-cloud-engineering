package config

import "github.com/Flo0807/hsfl-master-ai-cloud-engineering/bulletin-board-service/database"

type Config struct {
	HttpServerPort     string              `env:"HTTP_SERVER_PORT" envDefault:"3000"`
	GrpcServerPort     string              `env:"GRPC_SERVER_PORT" envDefault:"50052"`
	AuthServiceUrlGrpc string              `env:"AUTH_SERVICE_URL_GRPC"`
	Database           database.PsqlConfig `envPrefix:"DB_"`
}
