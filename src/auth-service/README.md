# Auth Service

The auth service is responsible for managing users. It provides a REST API to register and login users and a gRPC API to validate JWTs.

## Prerequisites

The auth service requires a running PostgreSQL database.

You can start a PostgreSQL database using Docker:

```bash
docker run --name postgres -e POSTGRES_PASSWORD=password -p 5432:5432 -d postgres
```

The auth service uses a ECDSA private key to sign JWTs. The key can be generated using the following command:

```bash
ssh-keygen -t ecdsa -f /path/to/key -m pem
```

Make sure to set the `JWT_PRIVATE_KEY` environment variable to the path of the generated key.

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
| GRPC_SERVER_PORT | Port the grpc server listens on | 50051 |
| JWT_PRIVATE_KEY | Path to ECDSA private key used to sign JWTs | - |
| DB_HOST | Hostname of PostgreSQL database | - |
| DB_PORT | Port of PostgreSQL database | - |
| DB_USER | Username of PostgreSQL database | - |
| DB_PASSWORD | Password of PostgreSQL database | - |
| DB_NAME | Name of PostgreSQL database | - |

You may also use a `.env` file to set environment variables. See `.env.example` for an example.
