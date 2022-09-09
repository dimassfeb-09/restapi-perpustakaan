package main

import (
	"github.com/dimassfeb-09/restapi-perpustakaan/app"
	"github.com/dimassfeb-09/restapi-perpustakaan/controller"
	"github.com/dimassfeb-09/restapi-perpustakaan/exception"
	"github.com/dimassfeb-09/restapi-perpustakaan/helper"
	"github.com/dimassfeb-09/restapi-perpustakaan/repository"
	"github.com/dimassfeb-09/restapi-perpustakaan/service"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.ForceConsoleColor()

	db := app.NewDB()
	r := gin.New()
	r.Use(gin.CustomRecovery(exception.ErrorHandler))
	r.HandleMethodNotAllowed = true

	userRepository := repository.NewUserRepositoryImpl()
	userService := service.NewUserServiceImpl(db, userRepository)
	userController := controller.NewUserControllerImpl(userService)

	r.POST("/user/add", userController.Create)
	r.PUT("/user/update/:id", userController.Update)
	r.DELETE("/user/delete/:id", userController.Delete)
	r.GET("/user/:id", userController.FindById)

	err := r.Run()
	helper.PanicIfError(err)
}
