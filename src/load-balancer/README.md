# Load Balancer

This is a simple load balancer that distributes requests to multiple servers.

It supports the following load balancing algorithms:

- Round Robin
- Random
- IP Hash

## Run

The load balancer only works within Docker, because it uses Docker networks to connect to the servers and IP addresses to identify the servers. You can start the load balancer using the provided `compose.yaml` in the root directory.

## Config

The Load Balancer can be configured by environment variables and a `config.json` file.

The following environment variables are available:

| Name | Description | Default |
| ---- | ----------- | -------- |
| HTTP_SERVER_PORT | Port the http server listens on | 3000 |

The default `config.json` looks like this:

```json
{
  "image": "docker.io/yeasy/simple-web",
  "networkName": "bulletinboard",
  "replicas": 3,
  "healthCheckIntervalSeconds": 5
}
```

The `image` field specifies the image that is used to start the servers. The `networkName` field specifies the name of the Docker network that is used to connect the servers to the load balancer. The `replicas` field specifies the number of servers that are started. The `healthCheckIntervalSeconds` field specifies the interval in which the load balancer checks the health of the servers.

You may also use a `.env` file to set environment variables. See `.env.example` for an example.