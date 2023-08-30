package repository

import (
	"Documentos/github.com/CelsodvJunior/CRUD-GO/src/configuration/logger"
	"Documentos/github.com/CelsodvJunior/CRUD-GO/src/configuration/rest_err"
	"Documentos/github.com/CelsodvJunior/CRUD-GO/src/model"
	"context"
	"os"
)
const (
	MONGODB_USER_DB= "MONGODB_USER_DB"
)
func (ur *userRepository) CreateUser(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init createUser repository")
	collection_name := os.Getenv(MONGODB_USER_DB)

	collection := ur.databaseConnection.Collection(collection_name)

	value, err := userDomain.GetJSONValue();
	 if err != nil {
		return nil, rest_err.NewInternalServerError(err.Error())
	}
	result, err := collection.InsertOne(context.Background(), value)
	if err != nil {
		return nil, rest_err.NewInternalServerError(err.Error())
	}
	userDomain.SetID(result.InsertedID.(string))

	return userDomain, nil
}
