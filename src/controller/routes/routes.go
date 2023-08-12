package routes

import (
	"Documentos/github.com/CelsodvJunior/CRUD-GO/src/controller"

	"github.com/gin-gonic/gin"
)

func InitRoutes(
	r *gin.RouterGroup,
	userControler controller.UserControllerInterface) {

	r.GET("/getUserById/:userId", userControler.FindUserById)
	r.GET("/getUserByEmail/:userEmail", userControler.FindeUserByEmail)
	r.POST("/creatUse", userControler.CreateUser)
	r.PUT("/updateUser/:userId", userControler.UpdateUser)
	r.DELETE("/deleteUse/:userId", userControler.DeleteUser)
}
