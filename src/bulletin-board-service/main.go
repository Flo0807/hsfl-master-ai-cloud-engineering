package main

import (
	"fmt"
	"github.com/Flo0807/hsfl-master-ai-cloud-engineering/bulletin-board-service/api/handler"
	"github.com/Flo0807/hsfl-master-ai-cloud-engineering/bulletin-board-service/api/router"
	"github.com/Flo0807/hsfl-master-ai-cloud-engineering/bulletin-board-service/database"
	"github.com/Flo0807/hsfl-master-ai-cloud-engineering/bulletin-board-service/models"
	"github.com/Flo0807/hsfl-master-ai-cloud-engineering/bulletin-board-service/repository"
	"github.com/Flo0807/hsfl-master-ai-cloud-engineering/bulletin-board-service/service"
	"gopkg.in/yaml.v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

var db *gorm.DB

type ApplicationConfig struct {
	Database database.PsqlConfig `yaml:"database"`
}

func main() {
	config, err := LoadFromConfigFile("config.yml")

	if err != nil {
		log.Fatal("Failed to load config file: ", err)
	}

	db, err = gorm.Open(postgres.Open(config.Database.Dsn()), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	err = db.AutoMigrate(&models.Post{})
	if err != nil {
		log.Fatal("Failed to migrate database: ", err)
	}

	postRepo := repository.NewPostRepository(db)
	postService := service.NewPostService(postRepo)
	postHandler := handler.NewPostHandler(postService)

	r := router.NewRouter(postHandler)

	port := ":8080"
	fmt.Println("Server is running on port", port)
	log.Fatal(http.ListenAndServe(port, r))
}

func LoadFromConfigFile(path string) (*ApplicationConfig, error) {
	f, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Println("Failed to close file: ", err)
		}
	}(f)

	var config ApplicationConfig

	if err := yaml.NewDecoder(f).Decode(&config); err != nil {
		return nil, err
	}

	return &config, nil

}
