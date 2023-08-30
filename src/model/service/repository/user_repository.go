package repository

import (
	"Documentos/github.com/CelsodvJunior/CRUD-GO/src/configuration/rest_err"
	"Documentos/github.com/CelsodvJunior/CRUD-GO/src/model"

	"go.mongodb.org/mongo-driver/mongo"
)

func NewUserRepository(
	database *mongo.Database,
)UserRepository{
	return &userRepository{
		database,
	}
}

type userRepository struct{
	databaseConnection *mongo.Database
}

type UserRepository interface {
	CreateUser(
		userDomain model.UserDomainInterface,
	) (model.UserDomainInterface, *rest_err.RestErr)
}
