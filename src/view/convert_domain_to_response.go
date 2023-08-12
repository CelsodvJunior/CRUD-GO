package view

import (
	"Documentos/github.com/CelsodvJunior/CRUD-GO/src/controller/response"
	"Documentos/github.com/CelsodvJunior/CRUD-GO/src/model"
)

func ConvertDomainToResponse(
	userDomain model.UserDomainInterface,
) response.UserResponse {
	return response.UserResponse{
		ID:    "",
		Email: userDomain.GetEmail(),
		Name:  userDomain.GetName(),
		Age:   userDomain.GetAge(),
	}
}
