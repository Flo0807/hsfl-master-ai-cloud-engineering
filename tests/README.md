# Tests

## Prerequisites

Run the kubernetes cluster.

Clear the database. The command searches for the postgres pod and executes the import script on it.

```bash
cat tests/data/clear.sql | kubectl exec -i $(kubectl get pods --no-headers -o custom-columns=":metadata.name" | grep postgres) -- psql -U admin -d board-hub password
```

Import the test data. The command searches for the postgres pod and executes the import script on it.

```bash
cat tests/data/import.sql | kubectl exec -i $(kubectl get pods --no-headers -o custom-columns=":metadata.name" | grep postgres) -- psql -U admin -d board-hub password
```

Make sure you use the correct username, password and database name based on the kiubernetes configuration.

The test script generates 5 to 20 lorem ipsum words as the content of the posts. Overall, there are 100 posts generated.

## Running the tests

TODO

## Machine specs

- CPU:
- RAM:
- OS:





