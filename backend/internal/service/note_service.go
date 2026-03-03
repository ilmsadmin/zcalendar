package service

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"

	"zplus-calendar/internal/common"
	"zplus-calendar/internal/dto"
	"zplus-calendar/internal/repository"
)

// NoteService defines the interface for note business logic
type NoteService interface {
	Create(ctx context.Context, userID string, req dto.CreateNoteRequest) (*dto.NoteResponse, error)
	GetByID(ctx context.Context, id, userID string) (*dto.NoteResponse, error)
	List(ctx context.Context, userID string, params repository.NoteListParams) ([]dto.NoteResponse, int64, error)
	Update(ctx context.Context, id, userID string, req dto.UpdateNoteRequest) (*dto.NoteResponse, error)
	Delete(ctx context.Context, id, userID string) error
	GetByDate(ctx context.Context, userID, date string) ([]dto.NoteResponse, error)
	Search(ctx context.Context, userID, query string, limit, offset int64) ([]dto.NoteResponse, int64, error)
}

type noteService struct {
	noteRepo repository.NoteRepository
	logger   *zap.Logger
}

func NewNoteService(noteRepo repository.NoteRepository, logger *zap.Logger) NoteService {
	return &noteService{
		noteRepo: noteRepo,
		logger:   logger,
	}
}

func (s *noteService) Create(ctx context.Context, userID string, req dto.CreateNoteRequest) (*dto.NoteResponse, error) {
	doc := dto.NoteDocument{
		UserID:     userID,
		Title:      req.Title,
		Content:    req.Content,
		Date:       req.Date,
		CategoryID: req.CategoryID,
		Tags:       req.Tags,
		IsPinned:   req.IsPinned,
		Color:      req.Color,
		LunarDay:   req.LunarDay,
		LunarMonth: req.LunarMonth,
		LunarYear:  req.LunarYear,
	}

	if err := s.noteRepo.Create(ctx, &doc); err != nil {
		s.logger.Error("Failed to create note", zap.Error(err))
		return nil, common.NewInternalError("Không thể tạo ghi chú", err)
	}

	resp := dto.ToNoteResponse(&doc)
	return &resp, nil
}

func (s *noteService) GetByID(ctx context.Context, id, userID string) (*dto.NoteResponse, error) {
	doc, err := s.noteRepo.GetByID(ctx, id, userID)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, common.NewNotFoundError("Ghi chú không tồn tại")
		}
		s.logger.Error("Failed to get note", zap.Error(err))
		return nil, common.NewInternalError("Không thể lấy ghi chú", err)
	}

	resp := dto.ToNoteResponse(doc)
	return &resp, nil
}

func (s *noteService) List(ctx context.Context, userID string, params repository.NoteListParams) ([]dto.NoteResponse, int64, error) {
	docs, total, err := s.noteRepo.List(ctx, userID, params)
	if err != nil {
		s.logger.Error("Failed to list notes", zap.Error(err))
		return nil, 0, common.NewInternalError("Không thể lấy danh sách ghi chú", err)
	}

	return dto.ToNoteListResponse(docs), total, nil
}

func (s *noteService) Update(ctx context.Context, id, userID string, req dto.UpdateNoteRequest) (*dto.NoteResponse, error) {
	// Verify existence
	_, err := s.noteRepo.GetByID(ctx, id, userID)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, common.NewNotFoundError("Ghi chú không tồn tại")
		}
		return nil, common.NewInternalError("Không thể lấy ghi chú", err)
	}

	update := bson.M{}
	if req.Title != nil {
		update["title"] = *req.Title
	}
	if req.Content != nil {
		update["content"] = *req.Content
	}
	if req.Date != nil {
		update["date"] = *req.Date
	}
	if req.CategoryID != nil {
		update["category_id"] = *req.CategoryID
	}
	if req.Tags != nil {
		update["tags"] = *req.Tags
	}
	if req.IsPinned != nil {
		update["is_pinned"] = *req.IsPinned
	}
	if req.Color != nil {
		update["color"] = *req.Color
	}
	if req.LunarDay != nil {
		update["lunar_day"] = *req.LunarDay
	}
	if req.LunarMonth != nil {
		update["lunar_month"] = *req.LunarMonth
	}
	if req.LunarYear != nil {
		update["lunar_year"] = *req.LunarYear
	}

	if len(update) == 0 {
		return s.GetByID(ctx, id, userID)
	}

	if err := s.noteRepo.Update(ctx, id, userID, update); err != nil {
		s.logger.Error("Failed to update note", zap.Error(err))
		return nil, common.NewInternalError("Không thể cập nhật ghi chú", err)
	}

	return s.GetByID(ctx, id, userID)
}

func (s *noteService) Delete(ctx context.Context, id, userID string) error {
	_, err := s.noteRepo.GetByID(ctx, id, userID)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return common.NewNotFoundError("Ghi chú không tồn tại")
		}
		return common.NewInternalError("Không thể xoá ghi chú", err)
	}

	if err := s.noteRepo.Delete(ctx, id, userID); err != nil {
		s.logger.Error("Failed to delete note", zap.Error(err))
		return common.NewInternalError("Không thể xoá ghi chú", err)
	}
	return nil
}

func (s *noteService) GetByDate(ctx context.Context, userID, date string) ([]dto.NoteResponse, error) {
	docs, err := s.noteRepo.GetByDate(ctx, userID, date)
	if err != nil {
		s.logger.Error("Failed to get notes by date", zap.Error(err))
		return nil, common.NewInternalError("Không thể lấy ghi chú theo ngày", err)
	}
	return dto.ToNoteListResponse(docs), nil
}

func (s *noteService) Search(ctx context.Context, userID, query string, limit, offset int64) ([]dto.NoteResponse, int64, error) {
	docs, total, err := s.noteRepo.Search(ctx, userID, query, limit, offset)
	if err != nil {
		s.logger.Error("Failed to search notes", zap.Error(err))
		return nil, 0, common.NewInternalError("Không thể tìm kiếm ghi chú", err)
	}
	return dto.ToNoteListResponse(docs), total, nil
}
