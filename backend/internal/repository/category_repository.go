package repository

import (
	"github.com/google/uuid"
	"gorm.io/gorm"

	"zplus-calendar/internal/model"
)

// CategoryRepository defines the interface for category data access
type CategoryRepository interface {
	Create(category *model.Category) error
	GetByID(id, userID uuid.UUID) (*model.Category, error)
	List(userID uuid.UUID, catType string) ([]model.Category, error)
	Update(category *model.Category) error
	Delete(id, userID uuid.UUID) error
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db: db}
}

func (r *categoryRepository) Create(category *model.Category) error {
	return r.db.Create(category).Error
}

func (r *categoryRepository) GetByID(id, userID uuid.UUID) (*model.Category, error) {
	var category model.Category
	err := r.db.Where("id = ? AND user_id = ?", id, userID).First(&category).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *categoryRepository) List(userID uuid.UUID, catType string) ([]model.Category, error) {
	var categories []model.Category
	query := r.db.Where("user_id = ?", userID)
	if catType != "" {
		query = query.Where("type = ? OR type = ?", catType, "both")
	}
	err := query.Order("sort_order ASC, name ASC").Find(&categories).Error
	return categories, err
}

func (r *categoryRepository) Update(category *model.Category) error {
	return r.db.Save(category).Error
}

func (r *categoryRepository) Delete(id, userID uuid.UUID) error {
	return r.db.Where("id = ? AND user_id = ?", id, userID).Delete(&model.Category{}).Error
}

// TagRepository defines the interface for tag data access
type TagRepository interface {
	Create(tag *model.Tag) error
	GetByID(id, userID uuid.UUID) (*model.Tag, error)
	List(userID uuid.UUID) ([]model.Tag, error)
	Update(tag *model.Tag) error
	Delete(id, userID uuid.UUID) error
}

type tagRepository struct {
	db *gorm.DB
}

func NewTagRepository(db *gorm.DB) TagRepository {
	return &tagRepository{db: db}
}

func (r *tagRepository) Create(tag *model.Tag) error {
	return r.db.Create(tag).Error
}

func (r *tagRepository) GetByID(id, userID uuid.UUID) (*model.Tag, error) {
	var tag model.Tag
	err := r.db.Where("id = ? AND user_id = ?", id, userID).First(&tag).Error
	if err != nil {
		return nil, err
	}
	return &tag, nil
}

func (r *tagRepository) List(userID uuid.UUID) ([]model.Tag, error) {
	var tags []model.Tag
	err := r.db.Where("user_id = ?", userID).Order("name ASC").Find(&tags).Error
	return tags, err
}

func (r *tagRepository) Update(tag *model.Tag) error {
	return r.db.Save(tag).Error
}

func (r *tagRepository) Delete(id, userID uuid.UUID) error {
	return r.db.Where("id = ? AND user_id = ?", id, userID).Delete(&model.Tag{}).Error
}
