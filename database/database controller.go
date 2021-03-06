// ---------------------------------------------------------------------------
// LN-ELECTRINIC PROJECT LN-16C10R
// wwww.ln-electronic.com  ProjectManager : @JJOY, @Kimera
// FileData: 29/3/2564 11:33
// ---------------------------------------------------------------------------
package database

import (
	"context"
	"encoding/hex"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/scrypt"
	"log"
	"loso/models"
	"loso/models/apperrors"
	"strconv"
	"strings"
)

// InserUser creates new user.
func (ln *LnDatabase) InsertUser(user *models.User) (*models.User, error) {
	// Specifies the order in which to return results.
	log.Println("binsert:", user.Passwd)
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

// FindByUser returns the username.
func (ln *LnDatabase) FindByUser(username string) (*models.User, error) {
	var user *models.User
	err := ln.DB.Collection("test").
		FindOne(context.Background(), bson.D{{Key: "username", Value: username}}).
		Decode(&user)
	if err != nil {
		log.Printf("Unable to get user with username : %v. Err: %v\n", username, err)
		return nil, apperrors.NewNotFound("username", username)
		//return nil, err

	}
	return user, err
}

//Check User Login
func (ln *LnDatabase) CheckLogin(ctx context.Context, u *models.User) error {
	uGet, err := ln.FindByUser(u.Username)
	//  return  details err
	if err != nil {
		return apperrors.NewAuthorization("username or passowrd is incorrect.")
	}

	// verify password - we previously created this method
	match, err := comparePassword(uGet.Passwd, u.Passwd)

	if err != nil {
		return apperrors.NewInternal()
	}

	if !match {
		return apperrors.NewAuthorization("Invalid Username or password.")
	}

	*u = *uGet
	return nil
}

func comparePassword(storedPassword string, suppliedPassword string) (bool, error) {
	pwsalt := strings.Split(storedPassword, ".")

	// check supplied password by kimera v1.3 @SHA256
	salt, err := hex.DecodeString(pwsalt[1])
	if err != nil {
		return false, fmt.Errorf("Unable to verify user password")
	}
	shash, err := scrypt.Key([]byte(suppliedPassword), salt, 32768, 8, 1, 32)

	return hex.EncodeToString(shash) == pwsalt[0], nil
}
