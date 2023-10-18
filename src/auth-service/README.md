# Auth Service

The auth service is responsible for managing users. It provides a REST API to create, read, update and delete users.

## Database

The auth service uses a PostgreSQL database to store user data.

Setup postgres database using docker:

```bash
docker run --name postgres -e POSTGRES_PASSWORD=password -p 5432:5432 -d postgres
```

## Config

Example config:

```yaml
jwt:
    secret: secret
    access_token:
        expiration: 3600
db:
    host: localhost
    port: 5432
    user: postgres
    password: postgres
    database: postgres
```