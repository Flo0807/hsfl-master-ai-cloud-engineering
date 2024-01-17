package repository

import (
	"context"
	"testing"

	"github.com/Flo0807/hsfl-master-ai-cloud-engineering/bulletin-board-service/database"
	"github.com/Flo0807/hsfl-master-ai-cloud-engineering/bulletin-board-service/models"
	"github.com/Flo0807/hsfl-master-ai-cloud-engineering/lib/containerhelpers"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestIntegrationPsqlRepository(t *testing.T) {
	postgresDb, err := containerhelpers.StartPostgres()

	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		postgresDb.Terminate(context.Background())
	})

	port, err := postgresDb.MappedPort(context.Background(), "5432")

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

	db, err := gorm.Open(postgres.Open(dbConfig.Dsn()), &gorm.Config{})

	db.AutoMigrate(&models.Post{})

	if err != nil {
		t.Fatal("Failed to connect to database: ", err)
	}

	repository := NewPostPsqlRepository(db)

	t.Cleanup(clearTables(t, repository.DB))

	t.Run("Create", func(t *testing.T) {
		t.Cleanup(clearTables(t, repository.DB))

		// given
		post := &models.Post{
			Title:   "Test Title",
			Content: "Test Description",
		}

		// when
		repository.Create(post)

		// then
		assert.NotNil(t, getPostFromDatabase(t, repository.DB, post.Title))
	})

	t.Run("FindAll", func(t *testing.T) {
		t.Cleanup(clearTables(t, repository.DB))

		// given
		post1 := &models.Post{
			Title:   "Test Title",
			Content: "Test Description",
		}

		post2 := &models.Post{
			Title:   "Test Title 2",
			Content: "Test Description 2",
		}

		insertPost(t, repository.DB, post1)
		insertPost(t, repository.DB, post2)

		// when
		posts := repository.FindAll(10, 0)

		// then
		assert.Equal(t, 2, len(posts.Records))
		assert.Equal(t, uint(1), posts.Page.CurrentPage)
		assert.Equal(t, uint(10), posts.Page.PageSize)
		assert.Equal(t, uint(2), posts.Page.TotalRecords)
		assert.Equal(t, uint(1), posts.Page.TotalPages)
	})

	t.Run("FindById", func(t *testing.T) {
		t.Cleanup(clearTables(t, repository.DB))

		// given
		post := &models.Post{
			Title:   "Test Title",
			Content: "Test Description",
		}

		insertPost(t, repository.DB, post)

		// when
		foundPost := repository.FindByID(post.ID)

		// then
		assert.NotNil(t, foundPost)
		assert.Equal(t, getPostFromDatabase(t, repository.DB, post.Title).ID, foundPost.ID)
	})

	t.Run("Update", func(t *testing.T) {
		t.Cleanup(clearTables(t, repository.DB))

		// given
		post := &models.Post{
			Title:   "Test Title",
			Content: "Test Description",
		}

		insertPost(t, repository.DB, post)

		// when
		post.Title = "Updated Title"
		repository.Update(post)

		// then
		assert.Equal(t, "Updated Title", getPostFromDatabase(t, repository.DB, post.Title).Title)
	})

	t.Run("Delete", func(t *testing.T) {
		t.Cleanup(clearTables(t, repository.DB))

		// given
		post := &models.Post{
			Title:   "Test Title",
			Content: "Test Description",
		}

		insertPost(t, repository.DB, post)
		insertedPost := getPostFromDatabase(t, repository.DB, post.Title)

		// when
		repository.Delete(insertedPost)

		var count int64
		db.Model(&models.Post{}).Where("title = ?", post.Title).Count(&count)

		// then
		assert.Equal(t, int64(0), count)
	})
}

func getPostFromDatabase(t *testing.T, db *gorm.DB, title string) *models.Post {
	var post models.Post
	db.First(&post, "title = ?", title)
	return &post
}

func clearTables(t *testing.T, db *gorm.DB) func() {
	return func() {
		db.Where("1 = 1").Delete(&models.Post{})
	}
}

func insertPost(t *testing.T, db *gorm.DB, post *models.Post) {
	id := db.Create(post)

	if id.Error != nil {
		t.Logf("could not insert post: %s", id.Error.Error())
		t.FailNow()
	}
}
