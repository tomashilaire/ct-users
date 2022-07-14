package domain

import "time"

type User struct {
	Id       string    `bson:"_id" json:"id"`
	Name     string    `bson:"name" json:"name"`
	LastName string    `bson:"last_name" json:"last_name"`
	Email    string    `bson:"email" json:"email"`
	Type     string    `bson:"type" json:"type"`
	Password string    `bson:"password" json:"password"`
	Created  time.Time `bson:"created" json:"created"`
	Updated  time.Time `bson:"updated" json:"updated"`
}

func NewUser(id string, name string, lastName string, email string, userType string, password string, created time.Time) *User {
	return &User{
		Id: id, Name: name, LastName: lastName, Email: email, Type: userType,
		Password: password, Created: created, Updated: created,
	}
}
