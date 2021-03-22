package database

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"loso/model"
	"testing"
)

func TestDatabaseSuite(t *testing.T) {
	suite.Run(t, new(DatabaseSuite))
}

type DatabaseSuite struct {
	suite.Suite
	db *LnDatabase
}

func (s *DatabaseSuite) BeforeTest(suiteName, testName string) {
	s.T().Log("--BeforeTest--")
	db, _ := NewCon("ln-smt")
	s.db = db
}

func (s *DatabaseSuite) AfterTest(suiteName, testName string) {
	s.db.Close()
}

func (s *DatabaseSuite) TestPost() {
	s.db.DB.Collection("test").Drop(nil)

	var err error
	for i := 1; i <= 25; i++ {
		// user1
		UserID, _ := primitive.ObjectIDFromHex("5c99bd941ba7b2304ad8c52a")
		article := (&model.Post{
			UserID: UserID,
			Title:  fmt.Sprintf("tile%d", i),
			Body:   "quia et suscipit\nsuscipit recusandae consequuntur expedita et cum\nreprehenderit molestiae ut ut quas totam\nnostrum rerum est autem sunt rem eveniet architecto",
		}).New()
		s.db.CreatePost(article)
	}
	assert.Nil(s.T(), err)
}

// CreatePost creates a post.
func (d *LnDatabase) CreatePost(post *model.Post) *model.Post {
	// Specifies the order in which to return results.
	upsert := true
	result := d.DB.Collection("test").
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
