package controller

import (
	"Documentos/github.com/CelsodvJunior/CRUD-GO/src/model/service"

	"github.com/gin-gonic/gin"
)

func NewUserControllerInterface(
	serviceInterface service.UserDomainService,
) UserControllerInterface {
	return &userControllerInterface{
		service: serviceInterface,
	}
}

type UserControllerInterface interface {
	FindUserById(c *gin.Context)
	FindeUserByEmail(c *gin.Context)

	DeleteUser(c *gin.Context)
	CreateUser(c *gin.Context)
	UpdateUser(c *gin.Context)
}

type userControllerInterface struct {
	service service.UserDomainService
}

// FindUserById implements UserControllerInterface.
func (*userControllerInterface) FindUserById(c *gin.Context) {
	panic("unimplemented")
}

// FindeUserByEmail implements UserControllerInterface.
func (*userControllerInterface) FindeUserByEmail(c *gin.Context) {
	panic("unimplemented")
}
