package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

// New Connect database
func NewCon(host, dbname string) (*LnDatabase, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(host))
	if err != nil {
		return nil, err
	}
	ctxping, cancel := context.WithTimeout(context.Background(), 6*time.Second)
	defer cancel()
	err = client.Ping(ctxping, readpref.Primary())
	if err != nil {
		return nil, err
	}
	db := client.Database(dbname)
	return &LnDatabase{DB: db, Client: client, Context: ctx}, nil
}

// Close Connection
func (ln *LnDatabase) Close() {
	ln.Client.Disconnect(ln.Context)
}

// Database is a wrapper
type LnDatabase struct {
	DB      *mongo.Database
	Client  *mongo.Client
	Context context.Context
}
