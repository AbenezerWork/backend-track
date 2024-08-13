package controllers

import (
	"fmt"
	"net/http"
	"task_managerv2/domain"
	"task_managerv2/infrastructure"
	"time"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserUsecase domain.UserUsecase
}

func InitUserController(usecase domain.UserUsecase) *UserController {
	return &UserController{UserUsecase: usecase}
}

func (uc *UserController) RegisterHandler(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid format"})
		return
	}
	hashedPassword, err := infrastructure.HashPassword(user.Password)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}
	user.Password = string(hashedPassword)
	id, err := uc.UserUsecase.AddUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func (uc *UserController) LoginHandler(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid format"})
		return
	}
	err := uc.UserUsecase.UserLogin(&user)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Could not authenticate user"})
		return
	}
	fmt.Println("----user", user)

	token, err := infrastructure.GenerateJWT(user.ID.Hex(), user.Role, user.Email, 24*time.Hour)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "User logged in successfully", "token": token})

}

func (uc *UserController) RemoveUser(c *gin.Context) {

	id := c.Param("id")

	err := uc.UserUsecase.RemoveUser(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "malformed id"})
	}
	c.Status(http.StatusNoContent)
}
