package repository

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"log"
	"task_managerv2/domain"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskRepository struct {
	nextID     int
	collection *mongo.Collection
}

func InitTaskRepository(database *mongo.Database) *TaskRepository {
	TaskManagerInstance := TaskRepository{}
	TaskManagerInstance.collection = database.Collection("tasks")
	return &TaskManagerInstance
}

func (tm *TaskRepository) AddTask(task *domain.Task) (primitive.ObjectID, error) {
	id, err := tm.collection.InsertOne(context.TODO(), *task)
	if err != nil {
		return primitive.ObjectID{}, err
	}
	return id.InsertedID.(primitive.ObjectID), nil
}

func (tm *TaskRepository) DeleteTask(id primitive.ObjectID) error {
	filter := bson.D{{"_id", id}}
	tm.collection.DeleteOne(context.TODO(), filter)
	return nil
}

func (tm *TaskRepository) GetTask(id primitive.ObjectID) (error, domain.Task) {
	var ret domain.Task
	err := tm.collection.FindOne(context.TODO(), bson.D{{"_id", id}}).Decode(&ret)
	if err != nil {
		fmt.Println(err)
	}
	return nil, ret
}

func (tm *TaskRepository) UpdateTask(id primitive.ObjectID, task domain.Task) error {
	log.Println(task)
	update := bson.D{}
	if task.Title != "" {
		update = append(update, bson.E{"title", task.Title})
	}
	if task.Description != "" {
		update = append(update, bson.E{"description", task.Description})
	}
	if task.Status != "" {
		update = append(update, bson.E{"status", task.Status})
	}
	if task.DueDate != (time.Time{}) {
		update = append(update, bson.E{"time", task.DueDate})
	}
	res, err := tm.collection.UpdateByID(context.TODO(), id, bson.D{{"$set", update}})
	if err != nil {
		return err
	}
	if res.MatchedCount == 0 {
		return errors.New("Couldn't find entry in database!")

	}
	return nil
}

func (tm *TaskRepository) AllTasks(id string) []domain.Task {
	var err error
	var cursor *mongo.Cursor
	if id == "" {
		cursor, err = tm.collection.Find(context.TODO(), bson.D{})
	} else {
		cursor, err = tm.collection.Find(context.TODO(), bson.D{{"userid", id}})
	}
	if err != nil {
		log.Fatal(err)
	}
	all := []domain.Task{}
	for cursor.Next(context.TODO()) {
		var elem domain.Task
		err := cursor.Decode(&elem)
		if err != nil {
			log.Fatal()
		}
		all = append(all, elem)

	}
	return all
}
