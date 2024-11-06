package router

import (
	"userProfileManagment/controller"

	"github.com/gin-gonic/gin"
)

// InitializeRouter sets up all the routes and their corresponding handlers
func InitializeRouter(r *gin.Engine, controller *controller.UserController) {

	r.POST("/createuser", controller.CreateUser)
	r.POST("/getuser", controller.GetUser)
	r.POST("/updateuser", controller.UpdateUser)
	r.POST("/deleteuser", controller.DeleteUser)
}
