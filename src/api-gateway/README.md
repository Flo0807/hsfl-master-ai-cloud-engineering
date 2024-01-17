# API Gateway

The API gateway is responsible for routing requests to the appropriate service.

## Run

```bash
go run main.go -config <path-to-config-file>
```

## Config

The API Gateway can be configured by environment variables and a `config.json` file.

The following environment variables are available:

| Name | Description | Default |
| ---- | ----------- | -------- |
| HTTP_SERVER_PORT | Port the http server listens on | 3000 |

The default `config.json` looks like this:

```json
{
    "services": [
        {
            "name": "frontend",
            "contextPath": "/",
            "targetURL": "http://localhost:3000"
        },
        {
            "name": "auth",
            "contextPath": "/auth",
            "targetURL": "http://localhost:3001"
        },
        {
            "name": "bulletin-board",
            "contextPath": "/bulletin-board",
            "targetURL": "http://localhost:3002"
        },
        {
            "name": "feed",
            "contextPath": "/feed",
            "targetURL": "http://localhost:3003"
        }
    ]
}
```

The `services` field specifies the services that are available. The `name` field specifies the name of the service. The `contextPath` field specifies the context path of the service. The `targetURL` field specifies the URL of the service.

You may also use a `.env` file to set environment variables. See `.env.example` for an example.