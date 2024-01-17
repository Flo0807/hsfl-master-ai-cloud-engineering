package user

import (
	"context"
	"database/sql"
	"testing"

	"github.com/Flo0807/hsfl-master-ai-cloud-engineering/lib/containerhelpers"
	"github.com/Flo0807/hsfl-master-ai-cloud-engineering/src/auth-service/database"
	"github.com/Flo0807/hsfl-master-ai-cloud-engineering/src/auth-service/pkg/model"
	"github.com/stretchr/testify/assert"
)

func TestIntegrationPsqlRepository(t *testing.T) {
	postgres, err := containerhelpers.StartPostgres()

	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		postgres.Terminate(context.Background())
	})

	port, err := postgres.MappedPort(context.Background(), "5432")

	if err != nil {
		t.Fatalf("could not get database container port: %s", err.Error())
	}

	dbConfig := database.PsqlConfig{
		Host:     "localhost",
		Port:     port.Int(),
		User:     "postgres",
		Password: "password",
		Name:     "postgres",
	}

	repository, err := NewPsqlRepository(dbConfig)

	if err != nil {
		t.Fatalf("could not create psql repository: %s", err.Error())
	}

	t.Cleanup(clearTables(t, repository.db))

	t.Run("Migrate", func(t *testing.T) {
		t.Cleanup(clearTables(t, repository.db))

		// when
		err := repository.Migrate()

		// then
		assert.NoError(t, err)
		assertTableExists(t, repository.db, "users", []string{"id", "email", "password"})
	})

	t.Run("FindUserByEmail", func(t *testing.T) {
		t.Cleanup(clearTables(t, repository.db))

		// given
		user := []model.DbUser{
			{
				Email:    "email@example.com",
				Password: []byte("password"),
			},
			{
				Email:    "user@example.com",
				Password: []byte("password"),
			},
		}

		for _, u := range user {
			insertUser(t, repository.db, &u)
		}

		// when
		id := getUserFromDatabase(t, repository.db, user[0].Email).ID
		foundUser, err := repository.FindUserByEmail(user[0].Email)

		// test
		assert.NoError(t, err)
		assert.Equal(t, id, foundUser.ID)
	})

	t.Run("CreateUser", func(t *testing.T) {
		t.Cleanup(clearTables(t, repository.db))

		// given
		user := model.DbUser{
			Email:    "email@example.com",
			Password: []byte("password"),
		}

		// when
		err := repository.CreateUser(&user)

		// then
		assert.NoError(t, err)
		assert.NotNil(t, getUserFromDatabase(t, repository.db, user.Email))
	})
}

func insertUser(t *testing.T, db *sql.DB, user *model.DbUser) {
	_, err := db.Exec(`insert into users (email, password) values ($1, $2)`, user.Email, user.Password)
	if err != nil {
		t.Logf("could not insert user: %s", err.Error())
		t.FailNow()
	}
}

func getUserFromDatabase(t *testing.T, db *sql.DB, email string) *model.DbUser {
	row := db.QueryRow(`select id from users where email = $1`, email)

	var user model.DbUser
	if err := row.Scan(&user.ID); err != nil {
		return nil
	}

	return &user
}

func clearTables(t *testing.T, db *sql.DB) func() {
	return func() {
		if _, err := db.Exec("delete from users"); err != nil {
			t.Logf("could not delete rows from users: %s", err.Error())
			t.FailNow()
		}
	}
}

func assertTableExists(t *testing.T, db *sql.DB, name string, columns []string) {
	rows, err := db.Query(`select column_name from information_schema.columns where table_name = $1`, name)
	if err != nil {
		t.Fail()
		return
	}

	scannedCols := make(map[string]struct{})
	for rows.Next() {
		var column string
		if err := rows.Scan(&column); err != nil {
			t.Logf("expected")
			t.FailNow()
		}

		scannedCols[column] = struct{}{}
	}

	if len(scannedCols) == 0 {
		t.Logf("expected table '%s' to exist, but not found", name)
		t.FailNow()
	}

	for _, col := range columns {
		if _, ok := scannedCols[col]; !ok {
			t.Logf("expected table '%s' to have column '%s'", name, col)
			t.Fail()
		}
	}
}
