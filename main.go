package main

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"loso/database"
)

func main() {
	CreatePost()

}

// CreatePost creates a post.
func (d *LnDatabase) CreatePost(post *model.Post) *model.Post {
	// Specifies the order in which to return results.
	upsert := true
	result := d.DB.Collection("user").
		FindOneAndReplace(context.Background(),
			bson.D{{Key: "_id", Value: post.ID}},
			post,
			&options.FindOneAndReplaceOptions{
				Upsert: &upsert,
			},
		)
	if result != nil {
		return post
	}
	return nil
}
