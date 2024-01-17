# Feed Service

The feed service is responsible for managing feeds. It provides a REST API to fetch different types of feeds.

## Prequisites

Create an `.env` file and adjust the values to your needs. See config section for more information. Make sure the bulletin board service is running.

## Run

```bash
go run main.go
```

You may also use the provided `compose.yaml` in the root directory to start all services.

## Config

Config values are set via environment variables. The following environment variables are available:

| Name | Description | Default |
| ---- | ----------- | -------- |
| HTTP_SERVER_PORT | Port the http server listens on | 3000 |
| BULLETIN_BOARD_SERVICE_URL_GRPC | URL of the bulletin board service gRPC endpoint | - |

You may also use a `.env` file to set environment variables. See `.env.example` for an example.