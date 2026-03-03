package database

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"zplus-calendar/internal/config"
)

type MongoDB struct {
	Client *mongo.Client
	DB     *mongo.Database
}

func NewMongo(cfg config.MongoConfig) (*MongoDB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOpts := options.Client().ApplyURI(cfg.URI)
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		return nil, fmt.Errorf("failed to connect mongo: %w", err)
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, fmt.Errorf("failed to ping mongo: %w", err)
	}

	db := client.Database(cfg.DBName)

	return &MongoDB{
		Client: client,
		DB:     db,
	}, nil
}

func (m *MongoDB) Disconnect() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	m.Client.Disconnect(ctx)
}

// Collection returns a MongoDB collection
func (m *MongoDB) Collection(name string) *mongo.Collection {
	return m.DB.Collection(name)
}
