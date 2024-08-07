package main

import (
	"github.com/joho/godotenv"
	"task_managerv2/data"
	"task_managerv2/router"
)

// TODO: GET /tasks: Get a list of all tasks.
// TODO: GET /tasks/:id: Get the details of a specific task.
// TODO: PUT /tasks/:id: Update a specific task. This endpoint should accept a JSON body with the new details of the task.
// TODO: DELETE /tasks/:id: Delete a specific task.
// TODO: POST /tasks: Create a new task. This endpoint should accept a JSON body with the task's title, description, due date, and status.

func main() {
	godotenv.Load()
	data.InitTaskManager()
	router.Route()

}
