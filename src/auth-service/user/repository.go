package user

import "github.com/Flo0807/hsfl-master-ai-cloud-engineering/src/auth-service/pkg/model"

type Repository interface {
	Migrate() error
	FindUserByName(username string) (*model.DbUser, error)
	CreateUser(user *model.DbUser) error
}
