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
				Username: "user",
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
				Username: "user",
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

	t.Run("findUserByName", func(t *testing.T) {
		t.Run("return error if query fails", func(t *testing.T) {
			// given
			name := "user"

			dbmock.
				ExpectQuery("SELECT id, username, password FROM users").
				WithArgs(name).
				WillReturnError(errors.New("database error"))

			// when
			_, err := repository.FindUserByName(name)

			// test
			assert.Error(t, err)
		})

		t.Run("return user if query succeeds", func(t *testing.T) {
			// given
			name := "user"

			dbmock.
				ExpectQuery("SELECT id, username, password FROM users").
				WithArgs(name).
				WillReturnRows(sqlmock.NewRows([]string{"id", "username", "password"}).AddRow(1, name, []byte("password")))

			// when
			user, err := repository.FindUserByName(name)

			// test
			assert.NoError(t, err)
			assert.Equal(t, 1, user.ID)
		})
	})
}
