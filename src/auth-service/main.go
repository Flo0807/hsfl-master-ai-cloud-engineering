package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Flo0807/hsfl-master-ai-cloud-engineering/src/auth-service/api/handler"
	"github.com/Flo0807/hsfl-master-ai-cloud-engineering/src/auth-service/api/router"
	"github.com/Flo0807/hsfl-master-ai-cloud-engineering/src/auth-service/auth"
	"github.com/Flo0807/hsfl-master-ai-cloud-engineering/src/auth-service/crypto"
	"github.com/Flo0807/hsfl-master-ai-cloud-engineering/src/auth-service/database"
	"github.com/Flo0807/hsfl-master-ai-cloud-engineering/src/auth-service/user"
	"gopkg.in/yaml.v3"
)

type ApplicationConfig struct {
	Jwt      auth.JwtConfig      `yaml:"jwt"`
	Database database.PsqlConfig `yaml:"database"`
}

func main() {
	port := flag.String("port", "8080", "port to listen on")
	flag.Parse()

	c, err := LoadFromConfigFile("config.yml")

	if err != nil {
		log.Fatalf("error while loading config: %s", err.Error())
	}

	userRepository, err := user.NewPsqlRepository(c.Database)

	if err != nil {
		log.Fatalf("error while creating user repository: %s", err.Error())
	}

	if err := userRepository.Migrate(); err != nil {
		log.Fatalf("could not migrate: %s", err.Error())
	}

	hasher := crypto.NewBcryptHasher()
	jwtTokenGenerator, err := auth.NewJwtTokenGenerator(c.Jwt)

	if err != nil {
		log.Fatalf("error while creating jwt token generator: %s", err.Error())
	}

	handler := router.NewRouter(
		handler.NewLoginHandler(userRepository, hasher, jwtTokenGenerator),
		handler.NewRegisterHandler(userRepository, hasher),
	)

	addr := fmt.Sprintf("0.0.0.0:%s", *port)
	if err := http.ListenAndServe(addr, handler); err != nil {
		log.Fatalf("error while listen and serve: %s", err.Error())
	}
}

func LoadFromConfigFile(path string) (*ApplicationConfig, error) {
	f, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer f.Close()

	var config ApplicationConfig

	if err := yaml.NewDecoder(f).Decode(&config); err != nil {
		return nil, err
	}

	return &config, nil

}
