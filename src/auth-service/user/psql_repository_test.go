package user

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Flo0807/hsfl-master-ai-cloud-engineering/src/auth-service/pkg/model"
	"github.com/stretchr/testify/assert"
)

func TestPsqlRepositoryTest(t *testing.T) {
	db, dbmock, err := sqlmock.New()

	if err != nil {
		t.Fatal(err)
	}

	repository := PsqlRepository{db}

	t.Run("createUser", func(t *testing.T) {
		t.Run("return error if query fails", func(t *testing.T) {
			// given
			user := &model.DbUser{
				Email:    "email",
				Password: []byte("password"),
			}

			dbmock.
				ExpectExec("INSERT INTO users").
				WillReturnError(errors.New("database error"))

			// when
			err := repository.CreateUser(user)

			// test
			assert.Error(t, err)
		})

		t.Run("return nil if query succeeds", func(t *testing.T) {
			//given
			user := &model.DbUser{
				Email:    "email",
				Password: []byte("password"),
			}

			dbmock.
				ExpectExec("INSERT INTO users").
				WillReturnResult(sqlmock.NewResult(1, 1))

			// when
			err := repository.CreateUser(user)

			// test
			assert.NoError(t, err)
		})
	})

	t.Run("findUserByEmail", func(t *testing.T) {
		t.Run("return error if query fails", func(t *testing.T) {
			// given
			email := "email@example.com"

			dbmock.
				ExpectQuery("SELECT id, email, password FROM users").
				WithArgs(email).
				WillReturnError(errors.New("database error"))

			// when
			_, err := repository.FindUserByEmail(email)

			// test
			assert.Error(t, err)
		})

		t.Run("return user if query succeeds", func(t *testing.T) {
			// given
			email := "email@example.com"

			dbmock.
				ExpectQuery("SELECT id, email, password FROM users").
				WithArgs(email).
				WillReturnRows(sqlmock.NewRows([]string{"id", "email", "password"}).AddRow(1, email, []byte("password")))

			// when
			user, err := repository.FindUserByEmail(email)

			// test
			assert.NoError(t, err)
			assert.Equal(t, 1, user.ID)
		})
	})
}
