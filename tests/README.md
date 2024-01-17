# Tests

We want to perform load tests against the request coalescing implementation to validate the performance of the implementation. The load tests are performed against a local kubernetes cluster running on Minikube (see [kubernetes documentation](../k8s/README.md) for more information on how to set up a local cluster).

The following describes how to run the load tests.

## Prerequisites

Run all commands from the root directory of the project.

Start the kubernetes cluster.

Clear the database. The following command searches for the postgres pod and executes the import script on it.

```bash
cat tests/data/clear.sql | kubectl exec -i $(kubectl get pods --no-headers -o custom-columns=":metadata.name" | grep postgres) -- psql -U admin -d board-hub password
```

Import the test data. The following command searches for the postgres pod and executes the import script on it.

```bash
cat tests/data/import.sql | kubectl exec -i $(kubectl get pods --no-headers -o custom-columns=":metadata.name" | grep postgres) -- psql -U admin -d board-hub password
```

Make sure you use the correct username, password and database name based on the kubernetes configuration.

The import script generates 5 to 20 lorem ipsum words as the content of the posts. Overall, there are 100 posts generated.

## Database Dump

To reproduce the load test with the same data a database dump is located in `tests/data/dump.sql`. The dump contains 100 posts with 5 to 20 lorem ipsum words as the content of the posts.

## Running the tests

We use [hey](https://github.com/rakyll/hey) with the following command to perform load tests against two endpoints.

```bash
hey -z 3m http://localhost/bulletin-board/posts
```

```bash
hey -z 3m http://localhost/bulletin-board/posts-request-coalescing
```

The commands above run the tesst for 3 minutes with a default of 50 concurrent requests. We test the following endpoints:
- `http://localhost/bulletin-board/posts/bulletin-board/posts`
- `http://localhost/bulletin-board/posts-request-coalescing/bulletin-board/posts-request-coalescing`

## Machine specs

The load test was perfomed against a local kubernetes cluster running on Minikube.

- Device: MacBook Pro 2021
- CPU: Apple M1 Pro
- RAM: 16GB
- OS: macOS Sonoma Version 14.2.1





