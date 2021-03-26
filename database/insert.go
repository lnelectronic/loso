package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"loso/models"
	"strconv"
)

// InserUser creates new user.
func (ln *LnDatabase) InsertUser(user *models.User) (*models.User, error) {
	// Specifies the order in which to return results.
	result, err := ln.DB.Collection("test").InsertOne(
		context.Background(),
		user,
	)
	fmt.Println("InserData: ", result)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// GetUserByName returns the username.
func (ln *LnDatabase) GetUserByName(name string) *models.User {
	var user *models.User
	err := ln.DB.Collection("test").
		FindOne(context.Background(), bson.D{{Key: "username", Value: name}}).
		Decode(&user)
	if err != nil {
		return nil
	}
	return user
}

// GetUsers returns all users.
// Filter start, end int, order, sort string
func (ln *LnDatabase) GetUsers(fil *models.Filter) []*models.User {
	users := []*models.User{}
	cursor, err := ln.DB.Collection("test").
		Find(context.Background(), bson.D{},
			&options.FindOptions{
				Skip:  fil.Skip,
				Sort:  bson.D{bson.E{Key: fil.SortKey, Value: fil.SortVal}},
				Limit: fil.Limit,
			})
	if err != nil {
		return nil
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		user := &models.User{}
		if err := cursor.Decode(user); err != nil {
			return nil
		}
		users = append(users, user)
	}

	return users
}

// CountUser returns user count
func (ln *LnDatabase) CountUser() string {
	total, err := ln.DB.Collection("test").CountDocuments(context.Background(), bson.D{{}}, &options.CountOptions{})
	if err != nil {
		return "0"
	}
	return strconv.Itoa(int(total))
}

// GetUserByIDs returns user id.
func (ln *LnDatabase) GetUserByIDs(ids []primitive.ObjectID) []*models.User {
	var users []*models.User
	cursor, err := ln.DB.Collection("test").
		Find(context.Background(), bson.D{{
			Key: "_id",
			Value: bson.D{{
				Key:   "$in",
				Value: ids,
			}},
		}})
	if err != nil {
		return nil
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		user := &models.User{}
		if err := cursor.Decode(user); err != nil {
			return nil
		}
		users = append(users, user)
	}

	return users
}
