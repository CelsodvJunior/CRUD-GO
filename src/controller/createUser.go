package controller

import (
	"Documentos/github.com/CelsodvJunior/CRUD-GO/src/configuration/logger"
	"Documentos/github.com/CelsodvJunior/CRUD-GO/src/configuration/rest_err/validation"
	"Documentos/github.com/CelsodvJunior/CRUD-GO/src/controller/model/request"
	"Documentos/github.com/CelsodvJunior/CRUD-GO/src/controller/model/request/response"
	"net/http"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser controller",
		zap.String("journey", "CreateUser"),
	)
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err,
			zap.String("journey", "CreateUser"),
		)
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	response := response.UserResponse{
		ID:    "test",
		Email: userRequest.Email,
		Name:  userRequest.Name,
		Age:   userRequest.Age,
	}

	logger.Info("User created sucessfully",
		zap.String("journey", "CreateUser"),
	)
	c.JSON(http.StatusOK, response)
}
