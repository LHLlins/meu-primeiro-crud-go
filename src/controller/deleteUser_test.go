package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/LHLlins/meu-primeiro-crud-go/tree/main/src/controller/model"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupDeleteUserTests() *gin.Engine {
	resetGlobalState()
	users = append(users, model.User{ID: "1", Name: "UserToDelete", Email: "delete@example.com"})
	users = append(users, model.User{ID: "2", Name: "UserToKeep", Email: "keep@example.com"})
	nextID = 3

	router := gin.Default()
	router.DELETE("/users/:userId", DeleteUser) // Ensure param matches controller
	return router
}

func TestDeleteUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := setupDeleteUserTests()

	t.Run("SuccessfulDelete", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodDelete, "/users/1", nil)
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Len(t, users, 1) // User "1" should be deleted
		assert.Equal(t, "2", users[0].ID) // User "2" should remain
	})

	t.Run("UserNotFound", func(t *testing.T) {
		// reset state again to ensure users slice is as expected for this test
		router := setupDeleteUserTests()
		// Attempt to delete user "1" which was deleted in the previous sub-test if state is not reset
		// For safety, let's try deleting a user that was never there.
		req, _ := http.NewRequest(http.MethodDelete, "/users/99", nil)
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusNotFound, rr.Code)
		assert.Len(t, users, 2) // Ensure no user was deleted
	})
}
