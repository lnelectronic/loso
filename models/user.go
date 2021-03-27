// ---------------------------------------------------------------------------
// LN-ELECTRINIC PROJECT LN-16C10R
// wwww.ln-electronic.com  ProjectManager : @JJOY, @Kimera
// FileData: 24/3/2564 21:28
// ---------------------------------------------------------------------------
package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// user insert
type User struct {
	ID          primitive.ObjectID `bson:"_id" json:"id"`
	Username    string             `bson:"username" json:"username" validate:"required"`
	Passwd      string             `bson:"passwd" json:"passwd" validate:"required"`
	Email       string             `bson:"email" json:"email" validate:"required,email"`
	Phone       int64              `bson:"phone" json:"phone" validate:"required"`
	AccessLevel []string           `bson:"accessLevel" json:"accessLevel" validate:"required"`
	DataJoin    primitive.DateTime
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
		DataJoin:    u.DataJoin,
	}
}

// models regis user
type Createuser struct {
	ID          primitive.ObjectID `bson:"_id" json:"id"`
	Username    string             `bson:"username" json:"username" validate:"required,get=4,lte=16"`
	Password    string             `bson:"password" json:"password" validate:"required,get=6,lte=12"`
	JoinDate    primitive.DateTime `bson:"join_date" json:"join_date" validate:"required"`
	ChangeDate  primitive.DateTime `bson:"change_date" json:"change_date" validate:"required"`
	Fullname    string             `bson:"fullname" json:"fullname" validate:"required"`
	Surname     string             `bson:"surname" json:"surname" validate:"required"`
	Nickname    string             `bson:"nickname" json:"nickname" validate:"required"`
	Gender      string             `bson:"gender" json:"gender" validate:"required"`
	Birthdate   primitive.DateTime `bson:"birthdate" json:"birthdate" validate:"required"`
	MobilePhone string             `bson:"mobile_phone" json:"mobile_phone" validate:"required"`
	Telephone   string             `bson:"telephone" json:"telephone" validate:"required"`
	Email       string             `bson:"email" json:"email" validate:"required,email"`
	AddressId   primitive.ObjectID `bson:"address_id" json:"address_id"`
}

// New is an instance Createuser
func (c *Createuser) New() *Createuser {
	return &Createuser{
		ID:          primitive.NewObjectID(),
		Username:    c.Username,
		Password:    c.Password,
		JoinDate:    c.JoinDate,
		ChangeDate:  c.ChangeDate,
		Fullname:    c.Fullname,
		Surname:     c.Surname,
		Nickname:    c.Nickname,
		Gender:      c.Gender,
		Birthdate:   c.Birthdate,
		MobilePhone: c.MobilePhone,
		Telephone:   c.Telephone,
		Email:       c.Email,
		AddressId:   primitive.NewObjectID(),
	}
}
