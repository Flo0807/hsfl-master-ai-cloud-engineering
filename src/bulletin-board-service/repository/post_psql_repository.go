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

type PostPsqlRepository struct {
	DB *gorm.DB
}

func NewPostPsqlRepository(db *gorm.DB) *PostPsqlRepository {
	return &PostPsqlRepository{DB: db}
}

func (r *PostPsqlRepository) Create(post *models.Post) {
	r.DB.Create(post)
}

func (r *PostPsqlRepository) FindAll(take int64, skip int64) PostPage {
	var posts []models.Post
	var totalRecords int64

	r.DB.Model(&models.Post{}).Count(&totalRecords)

	r.DB.Limit(int(take)).Offset(int(skip)).Find(&posts)

	totalPages := (totalRecords + take - 1) / take

	return NewPostPage(uint(skip/take)+1, uint(take), uint(totalRecords), uint(totalPages), posts)
}

func (r *PostPsqlRepository) FindByID(id uint) models.Post {
	var post models.Post
	r.DB.First(&post, id)
	return post
}

func (r *PostPsqlRepository) Update(post *models.Post) {
	r.DB.Save(post)
}

func (r *PostPsqlRepository) Delete(post *models.Post) {
	r.DB.Delete(post)
}
