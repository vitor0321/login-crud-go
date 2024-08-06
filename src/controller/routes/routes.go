package routes

import (
	"example.com/mod/src/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(
	r *gin.Engine,
	userController controller.UserControllerInterface,
) {

	r.GET("/getUserById/:userId", userController.FindUserById)
	r.GET("/getUserByEmail/:userEmail", userController.FindUserByEmail)
	r.POST("/createUser", userController.CreateUser)
	r.PUT("/updateUser/:userId", userController.UpdateUser)
	r.DELETE("/deleteUser/:userId", userController.DeleteUser)
}
