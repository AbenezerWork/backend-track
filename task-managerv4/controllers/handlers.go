package controllers

import (
	"fmt"
	"net/http"
	"task_managerv2/model"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func GetTasksHandler(c *gin.Context) {
	user_id := c.MustGet("claims").(jwt.MapClaims)["user_id"]
	rep := GetTasks(user_id.(string))
	c.JSON(http.StatusOK, rep)
}

func DeleteTaskHandler(c *gin.Context) {
	user_id := c.MustGet("claims").(jwt.MapClaims)["user_id"]
	id := c.Param("id")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	rep, err := GetTask(objID)
	if user_id != rep.UserID {
		c.JSON(http.StatusBadRequest, gin.H{"error": "You don't have access to the requested resource"})
		return
	}

	err = DeleteTask(objID)
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}
	c.Status(http.StatusNoContent)
}

func GetTaskHandler(c *gin.Context) {
	user_id := c.MustGet("claims").(jwt.MapClaims)["user_id"]
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	rep, err := GetTask(objID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if user_id != rep.UserID {
		c.JSON(http.StatusBadRequest, gin.H{"error": "You don't have access to the requested resource"})
		return

	}

	c.JSON(http.StatusOK, rep)
}

func AddTaskHandler(c *gin.Context) {
	user_id := c.MustGet("claims").(jwt.MapClaims)["user_id"]
	var newTask model.Task
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newTask.UserID = user_id.(string)
	id, err := AddTask(newTask)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "couldn't add task"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "task created successfully", "id": id})
}

func UpdateTaskHandler(c *gin.Context) {
	user_id := c.MustGet("claims").(jwt.MapClaims)["user_id"]
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var newTask model.Task
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	rep, err := GetTask(objID)
	if user_id != rep.UserID {
		c.JSON(http.StatusBadRequest, gin.H{"error": "You don't have access to the requested resource"})
		return
	}
	newTask.UserID = rep.UserID

	UpdateTask(objID, newTask)
	c.JSON(http.StatusOK, gin.H{"message": "task updated successfully", "id": id})

}

func RegisterHandler(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid format"})
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}
	user.Password = string(hashedPassword)
	id, err := addUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": id.Hex()})
}

func LoginHandler(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid format"})
		return
	}
	err := UserLogin(&user)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Could not authenticate user"})
		return
	}
	fmt.Println("----user", user)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
	})
	jwtSecret := []byte("secret")

	jwtToken, err := token.SignedString(jwtSecret)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "User logged in successfully", "token": jwtToken})

}
