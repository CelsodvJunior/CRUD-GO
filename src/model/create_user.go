package model

import (
	"Documentos/github.com/CelsodvJunior/CRUD-GO/src/configuration/logger"
	"Documentos/github.com/CelsodvJunior/CRUD-GO/src/configuration/rest_err"
	"fmt"

	"go.uber.org/zap"
)

func (ud *UserDomain) CreateUser() *rest_err.RestErr {

	logger.Info("Init createUser model", zap.String("journey", "createUser"))

	ud.EcryptPassword()

	fmt.Println(ud)

	return nil
}
