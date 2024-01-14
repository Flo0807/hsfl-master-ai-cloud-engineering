package models

import (
	"github.com/Flo0807/hsfl-master-ai-cloud-engineering/bulletin-board-service/models"
)

type Feed struct {
	Posts []models.Post `json:"posts"`
}
