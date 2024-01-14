package config

type Config struct {
	HttpServerPort              string `env:"HTTP_SERVER_PORT" envDefault:"3000"`
	BuleltinBoardServiceUrlGrpc string `env:"BULLETIN_BOARD_SERVICE_URL_GRPC" envDefault:"50052"`
}
