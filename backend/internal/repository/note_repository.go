package repository

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"zplus-calendar/internal/database"
	"zplus-calendar/internal/dto"
)

const noteCollection = "notes"

// NoteRepository defines the interface for note data access (MongoDB)
type NoteRepository interface {
	Create(ctx context.Context, note *dto.NoteDocument) error
	GetByID(ctx context.Context, id, userID string) (*dto.NoteDocument, error)
	Update(ctx context.Context, id, userID string, update bson.M) error
	Delete(ctx context.Context, id, userID string) error
	List(ctx context.Context, userID string, params NoteListParams) ([]dto.NoteDocument, int64, error)
	GetByDate(ctx context.Context, userID, date string) ([]dto.NoteDocument, error)
	Search(ctx context.Context, userID, query string, limit, offset int64) ([]dto.NoteDocument, int64, error)
	EnsureIndexes(ctx context.Context) error
}

// NoteListParams holds filter/pagination for listing notes
type NoteListParams struct {
	Limit     int64
	Offset    int64
	DateFrom  string
	DateTo    string
	Tag       string
	IsPinned  *bool
	SortBy    string
	SortOrder int // 1 = asc, -1 = desc
}

type noteRepository struct {
	coll *mongo.Collection
}

func NewNoteRepository(mongoDB *database.MongoDB) NoteRepository {
	return &noteRepository{
		coll: mongoDB.Collection(noteCollection),
	}
}

func (r *noteRepository) Create(ctx context.Context, note *dto.NoteDocument) error {
	note.CreatedAt = time.Now()
	note.UpdatedAt = time.Now()
	if note.Tags == nil {
		note.Tags = []string{}
	}
	if note.Color == "" {
		note.Color = "#ffffff"
	}

	result, err := r.coll.InsertOne(ctx, note)
	if err != nil {
		return err
	}

	// Set the generated ID
	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		note.ID = oid.Hex()
	}
	return nil
}

func (r *noteRepository) GetByID(ctx context.Context, id, userID string) (*dto.NoteDocument, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var note dto.NoteDocument
	err = r.coll.FindOne(ctx, bson.M{
		"_id":     objID,
		"user_id": userID,
	}).Decode(&note)
	if err != nil {
		return nil, err
	}

	note.ID = id
	return &note, nil
}

func (r *noteRepository) Update(ctx context.Context, id, userID string, update bson.M) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	update["updated_at"] = time.Now()

	_, err = r.coll.UpdateOne(ctx,
		bson.M{"_id": objID, "user_id": userID},
		bson.M{"$set": update},
	)
	return err
}

func (r *noteRepository) Delete(ctx context.Context, id, userID string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = r.coll.DeleteOne(ctx, bson.M{
		"_id":     objID,
		"user_id": userID,
	})
	return err
}

func (r *noteRepository) List(ctx context.Context, userID string, params NoteListParams) ([]dto.NoteDocument, int64, error) {
	filter := bson.M{"user_id": userID}

	if params.DateFrom != "" && params.DateTo != "" {
		filter["date"] = bson.M{"$gte": params.DateFrom, "$lte": params.DateTo}
	} else if params.DateFrom != "" {
		filter["date"] = bson.M{"$gte": params.DateFrom}
	} else if params.DateTo != "" {
		filter["date"] = bson.M{"$lte": params.DateTo}
	}

	if params.Tag != "" {
		filter["tags"] = params.Tag
	}

	if params.IsPinned != nil {
		filter["is_pinned"] = *params.IsPinned
	}

	// Count total
	total, err := r.coll.CountDocuments(ctx, filter)
	if err != nil {
		return nil, 0, err
	}

	// Sort
	sortBy := params.SortBy
	if sortBy == "" {
		sortBy = "created_at"
	}
	sortOrder := params.SortOrder
	if sortOrder == 0 {
		sortOrder = -1
	}

	findOpts := options.Find().
		SetSort(bson.D{{Key: "is_pinned", Value: -1}, {Key: sortBy, Value: sortOrder}}).
		SetSkip(params.Offset).
		SetLimit(params.Limit)

	cursor, err := r.coll.Find(ctx, filter, findOpts)
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(ctx)

	var notes []dto.NoteDocument
	if err := cursor.All(ctx, &notes); err != nil {
		return nil, 0, err
	}

	// Fix IDs
	for i := range notes {
		if notes[i].ID == "" {
			// The _id was decoded as ObjectID, need to handle
		}
	}

	return notes, total, nil
}

func (r *noteRepository) GetByDate(ctx context.Context, userID, date string) ([]dto.NoteDocument, error) {
	filter := bson.M{"user_id": userID, "date": date}

	findOpts := options.Find().
		SetSort(bson.D{{Key: "is_pinned", Value: -1}, {Key: "created_at", Value: -1}})

	cursor, err := r.coll.Find(ctx, filter, findOpts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var notes []dto.NoteDocument
	if err := cursor.All(ctx, &notes); err != nil {
		return nil, err
	}

	return notes, nil
}

func (r *noteRepository) Search(ctx context.Context, userID, query string, limit, offset int64) ([]dto.NoteDocument, int64, error) {
	filter := bson.M{
		"user_id": userID,
		"$or": bson.A{
			bson.M{"title": bson.M{"$regex": query, "$options": "i"}},
			bson.M{"content": bson.M{"$regex": query, "$options": "i"}},
		},
	}

	total, err := r.coll.CountDocuments(ctx, filter)
	if err != nil {
		return nil, 0, err
	}

	findOpts := options.Find().
		SetSort(bson.D{{Key: "created_at", Value: -1}}).
		SetSkip(offset).
		SetLimit(limit)

	cursor, err := r.coll.Find(ctx, filter, findOpts)
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(ctx)

	var notes []dto.NoteDocument
	if err := cursor.All(ctx, &notes); err != nil {
		return nil, 0, err
	}

	return notes, total, nil
}

func (r *noteRepository) EnsureIndexes(ctx context.Context) error {
	indexes := []mongo.IndexModel{
		{
			Keys: bson.D{{Key: "user_id", Value: 1}, {Key: "date", Value: -1}},
		},
		{
			Keys: bson.D{{Key: "user_id", Value: 1}, {Key: "created_at", Value: -1}},
		},
		{
			Keys: bson.D{{Key: "user_id", Value: 1}, {Key: "tags", Value: 1}},
		},
		{
			Keys:    bson.D{{Key: "title", Value: "text"}, {Key: "content", Value: "text"}},
			Options: options.Index().SetDefaultLanguage("none"),
		},
	}

	_, err := r.coll.Indexes().CreateMany(ctx, indexes)
	return err
}
