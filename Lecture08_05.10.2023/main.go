package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func main() {
	r := gin.Default()

	r.Use(authMiddleware)

	userRoutes := r.Group("/users")
	{
		userRoutes.POST("/", createUser)
		userRoutes.GET("/", getUsers)
		userRoutes.GET("/:id", getUser)
		userRoutes.PUT("/:id", updateUser)
		userRoutes.DELETE("/:id", deleteUser)
	}

	r.Run(":8080")
}

// Хардкодированный токен для проверки авторизации
const authToken = "yourAuthToken"

func authMiddleware(c *gin.Context) {
	token := c.GetHeader("Authorization")

	if token != authToken {
		c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}

	c.Next()
}

type User struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

var users []User

func createUser(c *gin.Context) {
	var user User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	users = append(users, user)
	c.JSON(http.StatusCreated, user)
}

func getUsers(c *gin.Context) {
	c.JSON(http.StatusOK, users)
}

func getUser(c *gin.Context) {
	id := c.Param("id")
	index, err := strconv.Atoi(id)
	if err != nil || index < 0 || index >= len(users) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invalid index"})
		return
	}

	val := users[index]

	if val != (User{}) {
		c.JSON(http.StatusOK, val)
		return
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
}

func updateUser(c *gin.Context) {
	id := c.Param("id")
	index, err := strconv.Atoi(id)
	if err != nil || index < 0 || index >= len(users) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invalid index"})
		return
	}
	var updatedUser User

	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	users[index] = updatedUser
	c.JSON(http.StatusOK, updatedUser)
}

func deleteUser(c *gin.Context) {
	id := c.Param("id")
	index, err := strconv.Atoi(id)
	if err != nil || index < 0 || index >= len(users) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invalid index"})
		return
	}

	users = append(users[:index], users[index+1:]...)
	c.JSON(http.StatusNoContent, nil)
}
