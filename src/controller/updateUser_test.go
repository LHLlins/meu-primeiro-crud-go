package controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/LHLlins/meu-primeiro-crud-go/tree/main/src/controller/model"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupUpdateUserTests() *gin.Engine {
	resetGlobalState()
	users = append(users, model.User{ID: "1", Name: "Original Name", Email: "original@example.com"})
	nextID = 2

	router := gin.Default()
	router.PUT("/users/:userID", UpdateUser)
	return router
}

func TestUpdateUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := setupUpdateUserTests()

	t.Run("SuccessfulUpdate", func(t *testing.T) {
		updatedInfo := model.User{Name: "Updated Name", Email: "updated@example.com"}
		jsonValue, _ := json.Marshal(updatedInfo)
		req, _ := http.NewRequest(http.MethodPut, "/users/1", bytes.NewBuffer(jsonValue))
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		var updatedUser model.User
		err := json.Unmarshal(rr.Body.Bytes(), &updatedUser)
		assert.NoError(t, err)
		assert.Equal(t, "1", updatedUser.ID)
		assert.Equal(t, "Updated Name", updatedUser.Name)
		assert.Equal(t, "updated@example.com", updatedUser.Email)

		// Verify in-memory data
		assert.Equal(t, "Updated Name", users[0].Name)
	})

	t.Run("UserNotFound", func(t *testing.T) {
		updatedInfo := model.User{Name: "Updated Name", Email: "updated@example.com"}
		jsonValue, _ := json.Marshal(updatedInfo)
		req, _ := http.NewRequest(http.MethodPut, "/users/99", bytes.NewBuffer(jsonValue))
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusNotFound, rr.Code)
	})

    t.Run("InvalidPayload", func(t *testing.T) {
		router := setupUpdateUserTests() // Reset state for this specific sub-test if needed

		req, _ := http.NewRequest(http.MethodPut, "/users/1", bytes.NewBufferString(`{"name": "Test"`)) // Malformed JSON
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})
}
