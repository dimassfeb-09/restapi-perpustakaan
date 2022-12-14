package service

import (
	"context"
	"github.com/dimassfeb-09/restapi-perpustakaan/model/web/user"
)

type UserService interface {
	Create(ctx context.Context, request user.UserCreateRequest) user.UserResponse
	Update(ctx context.Context, request user.UserUpdateRequest) user.UserResponse
	Delete(ctx context.Context, userId int)
	FindById(ctx context.Context, userId int) (user.UserResponse, error)
	FindAll(ctx context.Context) []user.UserResponse
	LoginAuth(ctx context.Context, userName string, passWord string) (user.UserLoginResponse, error)
}
