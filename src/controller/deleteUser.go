package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteUser(c *gin.Context) {
	userID := c.Param("userId") // Note: router.go uses "userId", not "userID"

	for i, user := range users {
		if user.ID == userID {
			// Remove user from slice
			users = append(users[:i], users[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
}
