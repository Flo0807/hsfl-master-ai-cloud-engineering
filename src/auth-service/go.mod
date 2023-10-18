module github.com/Flo0807/hsfl-master-ai-cloud-engineering/src/auth-service

go 1.21.2

require (
	github.com/DATA-DOG/go-sqlmock v1.5.0
	github.com/golang-jwt/jwt v3.2.2+incompatible
	github.com/lib/pq v1.10.9
	github.com/stretchr/testify v1.8.4
	go.uber.org/mock v0.3.0
	golang.org/x/crypto v0.14.0
	gopkg.in/yaml.v3 v3.0.1
)

require (
	github.com/Flo0807/hsfl-master-ai-cloud-engineering/lib v0.0.0-20231018143811-3c7d81c81e29
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
)

replace github.com/Flo0807/hsfl-master-ai-cloud-engineering/lib => ../../lib
