package app

import (
	"database/sql"
	"github.com/dimassfeb-09/restapi-perpustakaan/controller"
	"github.com/dimassfeb-09/restapi-perpustakaan/repository"
	"github.com/dimassfeb-09/restapi-perpustakaan/service"
)

func NewInitializedUser(db *sql.DB) controller.UserController {
	userRepository := repository.NewUserRepositoryImpl()
	userService := service.NewUserServiceImpl(db, userRepository)
	userController := controller.NewUserControllerImpl(userService)

	return userController
}
