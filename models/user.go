package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// user insert
type User struct {
	ID          primitive.ObjectID `bson:"_id" json:"id"`
	Username    string             `bson:"username" json:"username" validate:"required"`
	Passwd      string             `bson:"passwd" json:"passwd"`
	Email       string             `bson:"email" json:"email" validate:"required,email"`
	Phone       int64              `bson:"phone" json:"phone"`
	AccessLevel []string           `bson:"accessLevel" json:"accessLevel"`
}

// New is an instance
func (u *User) New() *User {
	return &User{
		ID:          primitive.NewObjectID(),
		Username:    u.Username,
		Passwd:      u.Passwd,
		Email:       u.Email,
		Phone:       u.Phone,
		AccessLevel: u.AccessLevel,
	}
}
