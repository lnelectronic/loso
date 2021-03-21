package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"loso/config"
	"loso/route"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	route.Addsum()
	fmt.Println(config.Hostmgo)

	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://lndb:lnteam@103.212.181.187:27017")

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

	fmt.Println("Connected to MongoDB!")

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.Run(config.ServerHost)

}
