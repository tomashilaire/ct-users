package usersmongorepo

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	ctx, cancel := context.WithTimeout(context.Background(), cfg.DbConnTimeOut()*time.Second)
	defer cancel()
	client, err := mongo.NewClient(options.Client().ApplyURI(cfg.Dsn()).SetTLSConfig(cfg.DbTlsConfig()))
	if err != nil {
		return nil, err
	}
	err = client.Connect(ctx)
	if err != nil {
		log.Fatalf("Failed to connect to cluster: %v", err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Failed to ping cluster: %v", err)
	}

	fmt.Println("Connected to DocumentDB!")
	return &conn{client: client, db: client.Database(cfg.DbName())}, nil
}

func (c *conn) Disconnect() {
	c.client.Disconnect(context.TODO())
}

func (c *conn) DB() *mongo.Database {
	return c.db
}
