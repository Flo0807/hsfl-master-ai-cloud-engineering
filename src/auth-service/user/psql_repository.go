package user

import (
	"database/sql"

	"github.com/Flo0807/hsfl-master-ai-cloud-engineering/src/auth-service/database"
	"github.com/Flo0807/hsfl-master-ai-cloud-engineering/src/auth-service/pkg/model"

	_ "github.com/lib/pq"
)

type PsqlRepository struct {
	db *sql.DB
}

func NewPsqlRepository(dbConfig database.PsqlConfig) (*PsqlRepository, error) {
	db, err := sql.Open("postgres", dbConfig.Dsn())

	if err != nil {
		return nil, err
	}

	return &PsqlRepository{db}, nil
}

const createUserTable = `
CREATE TABLE IF NOT EXISTS users (
	id SERIAL PRIMARY KEY,
	username VARCHAR(255) UNIQUE NOT NULL,
	password bytea NOT NULL
);
`

func (r *PsqlRepository) Migrate() error {
	_, err := r.db.Exec(createUserTable)

	return err
}

const findUserByNameQuery = `
SELECT id, username, password FROM users WHERE username = $1;
`

func (r *PsqlRepository) FindUserByName(username string) (*model.DbUser, error) {
	var user model.DbUser

	err := r.db.QueryRow(findUserByNameQuery, username).Scan(&user.ID, &user.Username, &user.Password)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

const createUserQuery = `
INSERT INTO users (username, password) VALUES ($1, $2);
`

func (r *PsqlRepository) CreateUser(user *model.DbUser) error {
	_, err := r.db.Exec(createUserQuery, user.Username, user.Password)

	return err
}
