package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// The Post holds test ln post price
type Post struct {
	ID     primitive.ObjectID `bson:"_id" json:"id"`
	UserID primitive.ObjectID `bson:"userId" json:"userId"`
	Title  string             `bson:"title" json:"title"`
	Body   string             `bson:"body" json:"body"`
}

// New is an instance
func (p *Post) New() *Post {
	return &Post{
		ID:     primitive.NewObjectID(),
		UserID: p.UserID,
		Title:  p.Title,
		Body:   p.Body,
	}
}

// Paging Model
type Paging struct {
	Skip      *int64
	Limit     *int64
	SortKey   string
	SortVal   int
	Condition interface{}
}
