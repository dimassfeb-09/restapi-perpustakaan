package service

import (
	"context"
	"database/sql"
	"github.com/dimassfeb-09/restapi-perpustakaan/exception"
	"github.com/dimassfeb-09/restapi-perpustakaan/helper"
	"github.com/dimassfeb-09/restapi-perpustakaan/model/domain"
	"github.com/dimassfeb-09/restapi-perpustakaan/model/web/user"
	"github.com/dimassfeb-09/restapi-perpustakaan/repository"
)

type UserServiceImpl struct {
	DB             *sql.DB
	UserRepository repository.UserRepository
}

func NewUserServiceImpl(DB *sql.DB, userRepository repository.UserRepository) UserService {
	return &UserServiceImpl{DB: DB, UserRepository: userRepository}
}

func (service *UserServiceImpl) Create(ctx context.Context, request user.UserCreateRequest) user.UserResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.RollbackOrCommit(tx)

	userUsername, _ := service.UserRepository.FindByUsername(ctx, tx, request.Username)
	if userUsername.Username == request.Username {
		panic(exception.NewErrorDuplicate("Username telah digunakan"))
	}

	userEmail, _ := service.UserRepository.FindByEmail(ctx, tx, request.Email)
	if userEmail.Email == request.Email {
		panic(exception.NewErrorDuplicate("Email telah digunakan"))
	}

	users := domain.User{
		Name:     request.Name,
		Username: request.Username,
		Password: request.Password,
		Email:    request.Email,
		Level:    request.Level,
		CreateAt: request.CreateAt,
	}

	userCreate := service.UserRepository.Create(ctx, tx, users)

	return helper.ToUserResponse(userCreate)
}

func (service *UserServiceImpl) Update(ctx context.Context, request user.UserUpdateRequest) user.UserResponse {

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.RollbackOrCommit(tx)

	_, err = service.UserRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewErrorNotFound(err.Error()))
	}

	users := domain.User{
		Name:     request.Name,
		Username: request.Username,
		Password: request.Password,
		Email:    request.Email,
		Level:    request.Level,
		CreateAt: request.CreateAt,
	}

	userUpdate := service.UserRepository.Update(ctx, tx, users)

	return helper.ToUserResponse(userUpdate)
}

func (service *UserServiceImpl) Delete(ctx context.Context, userId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.RollbackOrCommit(tx)

	findId, err := service.UserRepository.FindById(ctx, tx, userId)
	if err != nil {
		panic(exception.NewErrorNotFound(err.Error()))
	}

	service.UserRepository.Delete(ctx, tx, findId.Id)
}

func (service *UserServiceImpl) FindBy(ctx context.Context, filterBy string, value interface{}) user.UserResponse {

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	findBy, err := service.UserRepository.FindBy(ctx, tx, filterBy, value)
	if err != nil {
		panic(exception.NewErrorNotFound(err.Error()))
	}

	return helper.ToUserResponse(findBy)
}

func (service *UserServiceImpl) FindById(ctx context.Context, userId int) (user.UserResponse, error) {

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	findBy, err := service.UserRepository.FindById(ctx, tx, userId)
	if err != nil {
		panic(exception.NewErrorNotFound(err.Error()))
	}

	return helper.ToUserResponse(findBy), nil
}

func (service *UserServiceImpl) FindAll(ctx context.Context) []user.UserResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	users := service.UserRepository.FindAll(ctx, tx)

	return helper.ToUserResponses(users)
}
