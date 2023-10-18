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
	email VARCHAR(255) UNIQUE NOT NULL,
	password bytea NOT NULL
);
`

func (r *PsqlRepository) Migrate() error {
	_, err := r.db.Exec(createUserTable)

	return err
}

const findUserByEmailQuery = `
SELECT id, email, password FROM users WHERE email = $1;
`

func (r *PsqlRepository) FindUserByEmail(email string) (*model.DbUser, error) {
	var user model.DbUser

	err := r.db.QueryRow(findUserByEmailQuery, email).Scan(&user.ID, &user.Email, &user.Password)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

const createUserQuery = `
INSERT INTO users (email, password) VALUES ($1, $2);
`

func (r *PsqlRepository) CreateUser(user *model.DbUser) error {
	_, err := r.db.Exec(createUserQuery, user.Email, user.Password)

	return err
}
