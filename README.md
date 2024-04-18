#  BoardHub

[![](https://codecov.io/gh/Flo0807/hsfl-master-ai-cloud-engineering/graph/badge.svg?token=WILJH4U7EH)](https://codecov.io/gh/Flo0807/hsfl-master-ai-cloud-engineering)
![](https://github.com/Flo0807/hsfl-master-ai-cloud-engineering/actions/workflows/auth-service.yml/badge.svg)
![](https://github.com/Flo0807/hsfl-master-ai-cloud-engineering/actions/workflows/bulletin-board-service.yml/badge.svg)
![](https://github.com/Flo0807/hsfl-master-ai-cloud-engineering/actions/workflows/feed-service.yml/badge.svg)
![](https://github.com/Flo0807/hsfl-master-ai-cloud-engineering/actions/workflows/api-gateway.yml/badge.svg)

> [!NOTE]  
> This project is the submission for the course "Cloud Engineering" at [Flensburg University of Applied Sciences](https://hs-flensburg.de/) in the Master's program "Applied Computer Science". The goal of the course was to learn about cloud technologies and how to develop and deploy applications in the cloud. The task for the final project was to develop microservices in go and deploy them to a Kubernetes cluster.

BoardHub is an interactive platform that allows users to create and manage bulletin board posts. The platform offers an organized space for people to share news, ideas, or requests. Simple and easy to use, users can create bulletin board posts and comment on other posts. It features a user-friendly interface and different types of feeds to view posts. Whether it's for a business, a community, a club, a school, personal use or more, BoardHub provides a convenient way to communicate and connect.

## Use Cases

### University

BoardHub may be used to create a bulletin board for a university. Students can post about events, clubs, or other topics. The university can also post about important news, announcements or job opportunities and internships. Students can search for project partners or find supervisors for their thesis.

### Business

BoardHub can provide an ideal platform for businesses. Whether it's posting important updates, celebrating employee achievements or sharing project opportunities, BoardHub provides a space for company-wide communication.

### Private Community

For private communities like neighborhood associations or clubs, BoardHub serves as a common space to share updates, arrange events or discuss community concerns. Residents can post about local issues, lost or found items, neighborhood recommendations or any general inquiries.

## Setup

## Local Setup

### Prequisites

Docker and Docker Compose must be installed on your machine.

A ECDSA private key is required to sign JWTs. The key can be generated using the following command:

```bash
ssh-keygen -t ecdsa -f ./src/auth-service/key -m pem
```

When using the provided `compose.yaml` file, the key must be placed in the `auth-service` directory and named `key`.

You may configure the path to the key and the name using the `JWT_PRIVATE_KEY` environment variable (see `compose.yml`).

Build the frontend using the following commands in the `frontend` directory:

```bash
npm install
npm run build
```

### Starting the application

We provide a `compose.yaml` file to run the application. To start the application, run the following command in the root directory of the project:

```bash
docker compose up
```

### Accessing the application

The frontend is accessible at `http://localhost:3000`.

The default api gateway config uses the following paths to route requests to the corresponding services:

| Path              | Service       |
| ----------------- | ------------- |
| `/`               | frontend      |
| `/auth`           | auth-service  |
| `/bulletin-board` | bulletin-feed |
| `/feed`           | feed-service  |

### Example Data

When started via docker-compose the database is automatically populated with example data when the application is started for the first time. We provide two sql files located in the `scripts` directory. The `create.sql` file creates the database and the tables. The `insert.sql` file inserts example data into the database.

An example user is created with the following credentials:

| Username              | Password   |
| ------------------ | ---------- |
| `test` | `test` |

Along with some example posts.

## Kubernetes Setup

See [Kubernetes Setup](/k8s/README.md) for instructions on how to deploy the application to a Kubernetes cluster and how to import example data.

## Authors

Florian Arens\
florian.arens@stud.hs-flensburg.de\
Hochschule Flensburg

Tiark Millat\
tiark.millat@stud.hs-flensburg.de\
Hochschule Flensburg

Finn Wessel\
finn.wessel@stud.hs-flensburg.de\
Hochschule Flensburg