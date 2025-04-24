package routes

import (
	"github.com/LHLlins/meu-primeiro-crud-go/tree/main/src/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {
	r.GET("/getUserById/:userID", controller.FindUserById)
	r.GET("/getUserByEmail/:userEmail", controller.FindUserByEmail)
	r.POST("/createUser", controller.CreateUSer)
	r.PUT("/updateUser/:userID", controller.UpdateUser)
	r.DELETE("/deleteUser/:userId", controller.DeleteUser)
}
