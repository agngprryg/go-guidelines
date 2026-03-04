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

func (r *CategoryRepository) Create(category *models.Category)error{
	return r.db.Create(category).Error
}

func (r *CategoryRepository) Update(id uint, data map[string]interface{})(*models.Category, error){
	var category models.Category
	if err := r.db.First(&category, id).Error; err !=nil{
		return nil, err
	}

	if err := r.db.Model(&category).Updates(data).Error;err !=nil{
		return nil, err
	}

	return &category, nil
}

func (r *CategoryRepository) Delete(id uint) error {
	return r.db.Delete(&models.Category{}, id).Error
}