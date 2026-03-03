package service

import (
	"github.com/google/uuid"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"zplus-calendar/internal/common"
	"zplus-calendar/internal/dto"
	"zplus-calendar/internal/model"
	"zplus-calendar/internal/repository"
)

// ── Category Service ───────────────────────────────────────────────

type CategoryService interface {
	Create(userID uuid.UUID, req dto.CreateCategoryRequest) (*dto.CategoryResponse, error)
	List(userID uuid.UUID, catType string) ([]dto.CategoryResponse, error)
	Update(id, userID uuid.UUID, req dto.UpdateCategoryRequest) (*dto.CategoryResponse, error)
	Delete(id, userID uuid.UUID) error
}

type categoryService struct {
	repo   repository.CategoryRepository
	logger *zap.Logger
}

func NewCategoryService(repo repository.CategoryRepository, logger *zap.Logger) CategoryService {
	return &categoryService{repo: repo, logger: logger}
}

func (s *categoryService) Create(userID uuid.UUID, req dto.CreateCategoryRequest) (*dto.CategoryResponse, error) {
	category := model.Category{
		UserID:    userID,
		Name:      req.Name,
		Color:     req.Color,
		Icon:      req.Icon,
		Type:      req.Type,
		SortOrder: req.SortOrder,
	}

	if category.Color == "" {
		category.Color = "#3B82F6"
	}

	if err := s.repo.Create(&category); err != nil {
		s.logger.Error("Failed to create category", zap.Error(err))
		return nil, common.NewInternalError("Không thể tạo danh mục", err)
	}

	resp := dto.ToCategoryResponse(&category)
	return &resp, nil
}

func (s *categoryService) List(userID uuid.UUID, catType string) ([]dto.CategoryResponse, error) {
	categories, err := s.repo.List(userID, catType)
	if err != nil {
		s.logger.Error("Failed to list categories", zap.Error(err))
		return nil, common.NewInternalError("Không thể lấy danh sách danh mục", err)
	}
	return dto.ToCategoryListResponse(categories), nil
}

func (s *categoryService) Update(id, userID uuid.UUID, req dto.UpdateCategoryRequest) (*dto.CategoryResponse, error) {
	category, err := s.repo.GetByID(id, userID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.NewNotFoundError("Danh mục không tồn tại")
		}
		return nil, common.NewInternalError("Không thể lấy thông tin danh mục", err)
	}

	if req.Name != nil {
		category.Name = *req.Name
	}
	if req.Color != nil {
		category.Color = *req.Color
	}
	if req.Icon != nil {
		category.Icon = *req.Icon
	}
	if req.Type != nil {
		category.Type = *req.Type
	}
	if req.SortOrder != nil {
		category.SortOrder = *req.SortOrder
	}

	if err := s.repo.Update(category); err != nil {
		s.logger.Error("Failed to update category", zap.Error(err))
		return nil, common.NewInternalError("Không thể cập nhật danh mục", err)
	}

	resp := dto.ToCategoryResponse(category)
	return &resp, nil
}

func (s *categoryService) Delete(id, userID uuid.UUID) error {
	if _, err := s.repo.GetByID(id, userID); err != nil {
		if err == gorm.ErrRecordNotFound {
			return common.NewNotFoundError("Danh mục không tồn tại")
		}
		return common.NewInternalError("Không thể xoá danh mục", err)
	}

	if err := s.repo.Delete(id, userID); err != nil {
		s.logger.Error("Failed to delete category", zap.Error(err))
		return common.NewInternalError("Không thể xoá danh mục", err)
	}
	return nil
}

// ── Tag Service ────────────────────────────────────────────────────

type TagService interface {
	Create(userID uuid.UUID, req dto.CreateTagRequest) (*dto.TagResponse, error)
	List(userID uuid.UUID) ([]dto.TagResponse, error)
	Update(id, userID uuid.UUID, req dto.UpdateTagRequest) (*dto.TagResponse, error)
	Delete(id, userID uuid.UUID) error
}

type tagService struct {
	repo   repository.TagRepository
	logger *zap.Logger
}

func NewTagService(repo repository.TagRepository, logger *zap.Logger) TagService {
	return &tagService{repo: repo, logger: logger}
}

func (s *tagService) Create(userID uuid.UUID, req dto.CreateTagRequest) (*dto.TagResponse, error) {
	tag := model.Tag{
		UserID: userID,
		Name:   req.Name,
		Color:  req.Color,
	}

	if tag.Color == "" {
		tag.Color = "#6B7280"
	}

	if err := s.repo.Create(&tag); err != nil {
		s.logger.Error("Failed to create tag", zap.Error(err))
		return nil, common.NewInternalError("Không thể tạo nhãn", err)
	}

	resp := dto.ToTagResponse(&tag)
	return &resp, nil
}

func (s *tagService) List(userID uuid.UUID) ([]dto.TagResponse, error) {
	tags, err := s.repo.List(userID)
	if err != nil {
		s.logger.Error("Failed to list tags", zap.Error(err))
		return nil, common.NewInternalError("Không thể lấy danh sách nhãn", err)
	}
	return dto.ToTagListResponse(tags), nil
}

func (s *tagService) Update(id, userID uuid.UUID, req dto.UpdateTagRequest) (*dto.TagResponse, error) {
	tag, err := s.repo.GetByID(id, userID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.NewNotFoundError("Nhãn không tồn tại")
		}
		return nil, common.NewInternalError("Không thể lấy thông tin nhãn", err)
	}

	if req.Name != nil {
		tag.Name = *req.Name
	}
	if req.Color != nil {
		tag.Color = *req.Color
	}

	if err := s.repo.Update(tag); err != nil {
		s.logger.Error("Failed to update tag", zap.Error(err))
		return nil, common.NewInternalError("Không thể cập nhật nhãn", err)
	}

	resp := dto.ToTagResponse(tag)
	return &resp, nil
}

func (s *tagService) Delete(id, userID uuid.UUID) error {
	if _, err := s.repo.GetByID(id, userID); err != nil {
		if err == gorm.ErrRecordNotFound {
			return common.NewNotFoundError("Nhãn không tồn tại")
		}
		return common.NewInternalError("Không thể xoá nhãn", err)
	}

	if err := s.repo.Delete(id, userID); err != nil {
		s.logger.Error("Failed to delete tag", zap.Error(err))
		return common.NewInternalError("Không thể xoá nhãn", err)
	}
	return nil
}
