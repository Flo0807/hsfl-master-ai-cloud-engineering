# Bulletin Board Service

The Bulletin Board service provides a REST API to create, read, update and delete bulletin board data.

## Prerequisites

The bulletin board service requires a running PostgreSQL database and a running authentication service.

You can start a PostgreSQL database using Docker:

```bash
docker run --name postgres -e POSTGRES_PASSWORD=postgres -p 5432:5432 -d postgres
```

See authentication service documentation for instructions on how to start the authentication service.

Create an `.env` file and adjust the values to your needs. See config section for more information.

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
| GRPC_SERVER_PORT | Port the gRPC server listens on | 50051 |
| AUTH_SERVICE_URL_GRPC | URL of the authentication service gRPC endpoint | - |
| DB_HOST | Hostname of PostgreSQL database | - |
| DB_PORT | Port of PostgreSQL database | - |
| DB_USER | Username of PostgreSQL database | - |
| DB_PASSWORD | Password of PostgreSQL database | - |
| DB_NAME | Name of PostgreSQL database | - |

You may also use a `.env` file to set environment variables. See `.env.example` for an example.