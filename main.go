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
	initializerCategory := app.NewInitializedCategories(db)
	initializerBook := app.NewInitializedBook(db)
	initializedOfficer := app.NewInitializedOfficer(db)
	initializedGuestBook := app.NewInitializedGuestBook(db)

	r := app.NewRouter()
	app.NewUserRouter(r, initializedUser)
	app.NewCategoryRouter(r, initializerCategory)
	app.NewBookRouter(r, initializerBook)
	app.NewOfficerRouter(r, initializedOfficer)
	app.NewGuestBookRouter(r, initializedGuestBook)

	err := r.Run(":8080")
	helper.PanicIfError(err)
}
