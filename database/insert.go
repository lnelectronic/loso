package database

import (
	"context"
	"fmt"
	"loso/models"
)

// InserUser creates new user.
func (m *LnDatabase) InsertUser(user *models.User) (*models.User, error) {
	// Specifies the order in which to return results.
	result, err := m.DB.Collection("test").InsertOne(
		context.Background(),
		user,
	)
	fmt.Println(">>>>", result)
	if err != nil {
		return nil, err
	}
	return user, nil
}
