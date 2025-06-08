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

func TestCreateUSer(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("SuccessfulCreation", func(t *testing.T) {
		resetGlobalState()
		router := gin.Default()
		router.POST("/users", CreateUSer)

		user := model.User{Name: "Test User", Email: "test@example.com"}
		jsonValue, _ := json.Marshal(user)
		req, _ := http.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(jsonValue))
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusCreated, rr.Code)

		var createdUser model.User
		err := json.Unmarshal(rr.Body.Bytes(), &createdUser)
		assert.NoError(t, err)
		assert.Equal(t, "1", createdUser.ID)
		assert.Equal(t, user.Name, createdUser.Name)
		assert.Equal(t, user.Email, createdUser.Email)

		// Check if user was added to the slice
		assert.Len(t, users, 1)
		assert.Equal(t, users[0].ID, "1")
	})

	t.Run("InvalidPayload", func(t *testing.T) {
		resetGlobalState()
		router := gin.Default()
		router.POST("/users", CreateUSer)

		req, _ := http.NewRequest(http.MethodPost, "/users", bytes.NewBufferString(`{"name": "Test"`)) // Malformed JSON
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})
}
