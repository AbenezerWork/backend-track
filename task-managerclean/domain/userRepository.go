package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserRepository interface {
	AddUser(user *User) (primitive.ObjectID, error)
	DeleteUser(id primitive.ObjectID) error
	GetUser(id primitive.ObjectID) (error, User)
	UpdateUser(id primitive.ObjectID, user User) error
}
