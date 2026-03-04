package repositories

import (
	"mvc-gorm/models"

	"gorm.io/gorm"
)

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository{
	return &CategoryRepository{db: db}
}

func (r *CategoryRepository) GetAll()([]models.Category, error){
	var categories []models.Category
	result := r.db.Find(&categories)
	return categories, result.Error
}

func (r *CategoryRepository) GetByID(id uint)(*models.Category, error){
	var Category models.Category
	result := r.db.Preload("Products").First(&Category, id)
	return &Category, result.Error
}