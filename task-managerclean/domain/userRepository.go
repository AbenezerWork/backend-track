package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserRepository interface {
	AddUser(user *User) (primitive.ObjectID, error)
	DeleteUser(id primitive.ObjectID) error
	GetUserByEmail(string) (User, error)
}
