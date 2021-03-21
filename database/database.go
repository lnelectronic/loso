package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"loso/config"
	"time"
)

// New connect  for the mongo-go-driver set client.
func New(dbname string) (*LnDatabase, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.Hostmgo))
	if err != nil {
		return nil, err
	}
	ctxping, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = client.Ping(ctxping, readpref.Primary())
	if err != nil {
		return nil, err
	}
	db := client.Database(dbname)
	return &LnDatabase{DB: db, Client: client, Context: ctx}, nil
}

// Close closes the mongo-go-driver connection.
func (ln *LnDatabase) Close() {
	ln.Client.Disconnect(ln.Context)
}


// Database is a wrapper for the mongo-go-driver.
type LnDatabase struct {
	DB      *mongo.Database
	Client  *mongo.Client
	Context context.Context
}

