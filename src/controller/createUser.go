package controller

import (
	"Documentos/github.com/CelsodvJunior/CRUD-GO/src/configuration/rest_err"
	"Documentos/github.com/CelsodvJunior/CRUD-GO/src/controller/model/request"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	log.Println("Init CreateUser controller")
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		log.Printf("Error trying to marshal object, error=%s\n", err.Error())
		errRest := rest_err.NewBadRequestError("Some fields are incorrect")

		c.JSON(errRest.Code, errRest)
		return
	}

	fmt.Println(userRequest)
}
