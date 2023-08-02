package controller

import (
	"Documentos/github.com/CelsodvJunior/CRUD-GO/src/configuration/rest_err"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	err := rest_err.NewBadRequestError("You colled the rout worng!")
	c.JSON(err.Code, err)
}
