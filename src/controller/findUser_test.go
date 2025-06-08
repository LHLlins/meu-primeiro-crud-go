package controller

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/LHLlins/meu-primeiro-crud-go/tree/main/src/controller/model"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupFindUserTests() *gin.Engine {
	resetGlobalState()
	// Pre-populate users for find tests
	users = append(users, model.User{ID: "1", Name: "User One", Email: "one@example.com"})
	users = append(users, model.User{ID: "2", Name: "User Two", Email: "two@example.com"})
	nextID = 3

	router := gin.Default()
	router.GET("/users/id/:userID", FindUserById)
	router.GET("/users/email/:userEmail", FindUserByEmail)
	return router
}

func TestFindUserById(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := setupFindUserTests()

	t.Run("Found", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/users/id/1", nil)
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		var foundUser model.User
		err := json.Unmarshal(rr.Body.Bytes(), &foundUser)
		assert.NoError(t, err)
		assert.Equal(t, "1", foundUser.ID)
		assert.Equal(t, "User One", foundUser.Name)
	})

	t.Run("NotFound", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/users/id/99", nil)
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusNotFound, rr.Code)
	})
}

func TestFindUserByEmail(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := setupFindUserTests()

	t.Run("Found", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/users/email/two@example.com", nil)
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		var foundUser model.User
		err := json.Unmarshal(rr.Body.Bytes(), &foundUser)
		assert.NoError(t, err)
		assert.Equal(t, "2", foundUser.ID)
		assert.Equal(t, "two@example.com", foundUser.Email)
	})

	t.Run("NotFound", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/users/email/nonexistent@example.com", nil)
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusNotFound, rr.Code)
	})
}
