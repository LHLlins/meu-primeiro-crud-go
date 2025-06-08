package controller

import (
	"fmt"
	"net/http"

	"github.com/LHLlins/meu-primeiro-crud-go/tree/main/src/controller/model"
	"github.com/gin-gonic/gin"
)

func CreateUSer(c *gin.Context) {
	var userRequest model.User
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userRequest.ID = fmt.Sprintf("%d", nextID)
	nextID++

	users = append(users, userRequest)

	c.JSON(http.StatusCreated, userRequest)
}
