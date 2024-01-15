package repository

import (
	"github.com/Flo0807/hsfl-master-ai-cloud-engineering/bulletin-board-service/models"
	"gorm.io/gorm"
)

type Page struct {
	CurrentPage  uint `json:"currentPage"`
	PageSize     uint `json:"pageSize"`
	TotalRecords uint `json:"totalRecords"`
	TotalPages   uint `json:"totalPages"`
}

type PostPage struct {
	Page    Page          `json:"page"`
	Records []models.Post `json:"records"`
}

func NewPostPage(currentPage, pageSize, totalRecords, totalPages uint, records []models.Post) PostPage {
	return PostPage{
		Page: Page{
			CurrentPage:  currentPage,
			PageSize:     pageSize,
			TotalRecords: totalRecords,
			TotalPages:   totalPages,
		},
		Records: records,
	}
}

// PostPsqlRepository handles database operations for Post
type PostPsqlRepository struct {
	DB *gorm.DB
}

// NewPostPsqlRepository creates a new PostPsqlRepository
func NewPostPsqlRepository(db *gorm.DB) *PostPsqlRepository {
	return &PostPsqlRepository{DB: db}
}

// Create a new post
func (r *PostPsqlRepository) Create(post *models.Post) {
	r.DB.Create(post)
}

// FindAll posts
func (r *PostPsqlRepository) FindAll(take int64, skip int64) PostPage {
	var posts []models.Post
	var totalRecords int64

	// Ermittle die Gesamtzahl der Datensätze
	r.DB.Model(&models.Post{}).Count(&totalRecords)

	// Abfrage der Datensätze mit Paginierung
	r.DB.Limit(int(take)).Offset(int(skip)).Find(&posts)

	// Berechnung der Anzahl der Seiten
	totalPages := (totalRecords + take - 1) / take

	// Erstelle ein Page-Objekt, um die paginierten Daten und Metadaten zurückzugeben
	return NewPostPage(uint(skip/take)+1, uint(take), uint(totalRecords), uint(totalPages), posts)
}

// FindByID finds a post by ID
func (r *PostPsqlRepository) FindByID(id uint) models.Post {
	var post models.Post
	r.DB.First(&post, id)
	return post
}

// Update a post
func (r *PostPsqlRepository) Update(post *models.Post) {
	r.DB.Save(post)
}

// Delete a post
func (r *PostPsqlRepository) Delete(post *models.Post) {
	r.DB.Delete(post)
}

// Return the number of posts
func (r *PostPsqlRepository) Count() int64 {
	var count int64
	r.DB.Model(&models.Post{}).Count(&count)
	return count
}
