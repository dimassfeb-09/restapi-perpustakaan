package helper

import (
	"github.com/dimassfeb-09/restapi-perpustakaan/model/domain"
	"github.com/dimassfeb-09/restapi-perpustakaan/model/web/user"
)

func ToUserResponse(domainUser domain.User) user.UserResponse {
	return user.UserResponse{
		Id:       domainUser.Id,
		Name:     domainUser.Name,
		Username: domainUser.Username,
		Email:    domainUser.Email,
		Level:    domainUser.Level,
		CreateAt: domainUser.CreateAt,
	}
}

func ToUserResponses(domainUsers []domain.User) []user.UserResponse {
	var userResponses []user.UserResponse
	for _, data := range domainUsers {
		userResponses = append(userResponses, ToUserResponse(data))
	}

	return userResponses
}
