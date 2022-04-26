package usersmongorepo

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Connection interface {
	Disconnect()
	DB() *mongo.Database
}

type conn struct {
	client *mongo.Client
	db     *mongo.Database
}

func NewConnection(cfg Config) (*conn, error) {
	fmt.Printf("Database url: %s\n", cfg.Dsn())
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.Dsn()))
	if err != nil {
		return nil, err
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		println(err)
		return nil, err
	}
	return &conn{client: client, db: client.Database(cfg.DbName())}, nil
}

func (c *conn) Disconnect() {
	c.client.Disconnect(context.TODO())
}

func (c *conn) DB() *mongo.Database {
	return c.db
}
