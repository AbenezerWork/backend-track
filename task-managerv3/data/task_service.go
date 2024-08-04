package data

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"log"
	"task_managerv2/model"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TaskManager struct {
	tasks      map[string]model.Task
	nextID     int
	collection *mongo.Collection
}

// Initialize the TaskManagerInstance
var TaskManagerInstance = TaskManager{}

func InitTaskManager() *TaskManager {
	TaskManagerInstance = TaskManager{}
	TaskManagerInstance.tasks = make(map[string]model.Task)
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	TaskManagerInstance.collection = client.Database("task-manager").Collection("tasks")

	return &TaskManagerInstance
}

func (tm *TaskManager) AddTask(task *model.Task) primitive.ObjectID {
	id, err := tm.collection.InsertOne(context.TODO(), *task)
	if err != nil {
		log.Fatal(err)
	}
	return id.InsertedID.(primitive.ObjectID)
}

func (tm *TaskManager) DeleteTask(id primitive.ObjectID) error {
	filter := bson.D{{"_id", id}}
	tm.collection.DeleteOne(context.TODO(), filter)
	return nil
}

func (tm *TaskManager) GetTask(id primitive.ObjectID) (error, model.Task) {
	var ret model.Task
	err := tm.collection.FindOne(context.TODO(), bson.D{{"_id", id}}).Decode(&ret)
	if err != nil {
		fmt.Println(err)
	}
	return nil, ret
}

func (tm *TaskManager) UpdateTask(id primitive.ObjectID, task model.Task) error {
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

func (tm *TaskManager) AllTasks() []model.Task {
	cursor, err := tm.collection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	all := []model.Task{}
	for cursor.Next(context.TODO()) {
		var elem model.Task
		err := cursor.Decode(&elem)
		if err != nil {
			log.Fatal()
		}
		all = append(all, elem)

	}
	log.Println(tm.tasks)
	return all
}
