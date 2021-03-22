package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"loso/config"
	"time"
)

// New Connect
func NewCon(dbname string) (*LnDatabase, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.Hostmgo))
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

// Close
func (ln *LnDatabase) Close() {
	ln.Client.Disconnect(ln.Context)
}

// Database is a wrapper
type LnDatabase struct {
	DB      *mongo.Database
	Client  *mongo.Client
	Context context.Context
}

func Connection() *mongo.Client {
	// Set client options
	clientOptions := options.Client().ApplyURI(config.Hostmgo)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println("Connected to MongoDB!")

	return client
}
