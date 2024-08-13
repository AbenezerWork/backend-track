package repository

import (
	"context"

	"task_managerv2/domain"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	Collection *mongo.Collection
}

func InitUserRepository(database *mongo.Database) domain.UserRepository {
	userRepositoryInstance := UserRepository{}
	userRepositoryInstance.Collection = database.Collection("users")
	return &userRepositoryInstance
}

func (tm *UserRepository) DeleteUser(id primitive.ObjectID) error {
	filter := bson.D{{"_id", id}}
	tm.Collection.DeleteOne(context.TODO(), filter)
	return nil
}

func (tm *UserRepository) AddUser(user *domain.User) (primitive.ObjectID, error) {
	id, err := tm.Collection.InsertOne(context.TODO(), *user)
	if err != nil {
		return primitive.ObjectID{}, err
	}

	return id.InsertedID.(primitive.ObjectID), nil

}

func (tm *UserRepository) GetUserByEmail(email string) (domain.User, error) {
	filter := bson.D{{"email", email}}
	res := tm.Collection.FindOne(context.TODO(), filter)

	if res.Err() != nil {
		return domain.User{}, res.Err()
	}
	var user domain.User

	err := res.Decode(&user)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}
