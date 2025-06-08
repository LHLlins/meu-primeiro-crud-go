package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindUserById(c *gin.Context) {
	userID := c.Param("userID")

	for _, user := range users {
		if user.ID == userID {
			c.JSON(http.StatusOK, user)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
}

func FindUserByEmail(c *gin.Context) {
	userEmail := c.Param("userEmail")

	for _, user := range users {
		if user.Email == userEmail {
			c.JSON(http.StatusOK, user)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
}
