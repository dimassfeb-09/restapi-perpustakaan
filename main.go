package main

import (
	"github.com/dimassfeb-09/restapi-perpustakaan/app"
	"github.com/dimassfeb-09/restapi-perpustakaan/exception"
	"github.com/dimassfeb-09/restapi-perpustakaan/helper"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.ForceConsoleColor()

	db := app.NewDB()
	r := gin.New()
	r.Use(gin.CustomRecovery(exception.ErrorHandler))
	r.HandleMethodNotAllowed = true

	initializedUser := app.NewInitializedUser(db)
	app.NewUserRouter(r, initializedUser)

	err := r.Run()
	helper.PanicIfError(err)
}
