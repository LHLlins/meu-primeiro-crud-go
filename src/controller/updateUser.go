package controller

import (
	"net/http"

	"github.com/LHLlins/meu-primeiro-crud-go/tree/main/src/controller/model"
	"github.com/gin-gonic/gin"
)

func UpdateUser(c *gin.Context) {
	userID := c.Param("userID")
	var userRequest model.User

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, user := range users {
		if user.ID == userID {
			users[i].Name = userRequest.Name
			users[i].Email = userRequest.Email
			c.JSON(http.StatusOK, users[i])
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
}
