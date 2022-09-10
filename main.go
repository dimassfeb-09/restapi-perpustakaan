package main

import (
	"github.com/dimassfeb-09/restapi-perpustakaan/app"
	"github.com/dimassfeb-09/restapi-perpustakaan/helper"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.ForceConsoleColor()

	db := app.NewDB()
	initializedUser := app.NewInitializedUser(db)

	r := app.NewRouter()
	app.NewUserRouter(r, initializedUser)

	err := r.Run()
	helper.PanicIfError(err)
}
