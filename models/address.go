package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// models Address
type Address struct {
	ID            primitive.ObjectID `bson:"_id" json:"id"`
	HouseNo       string             `bson:"houseno" json:"houseno" validate:"required"`
	Alley         string             `bson:"alley" json:"alley" validate:"required"`
	Road          string             `bson:"road" json:"road" validate:"required"`
	SubDistrictID primitive.ObjectID `bson:"subdistrictid" json:"subdistrictid"`
}

// New instance Address
func (u *Address) New() *Address {
	return &Address{
		ID:            primitive.NewObjectID(),
		HouseNo:       u.HouseNo,
		Alley:         u.Alley,
		Road:          u.Road,
		SubDistrictID: primitive.NewObjectID(),
	}
}
