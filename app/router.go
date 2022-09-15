package app

import (
	"github.com/dimassfeb-09/restapi-perpustakaan/controller"
	"github.com/dimassfeb-09/restapi-perpustakaan/exception"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.CustomRecovery(exception.ErrorHandler))
	r.HandleMethodNotAllowed = true

	return r
}

func NewUserRouter(r *gin.Engine, userController controller.UserController) {
	r.POST("/user/add", userController.Create)
	r.PUT("/user/update/:id", userController.Update)
	r.DELETE("/user/delete/:id", userController.Delete)
	r.GET("/user/:id", userController.FindById)
	r.GET("/user", userController.FindAll)
	r.POST("/user/auth/login", userController.LoginAuth)
}

func NewCategoryRouter(r *gin.Engine, categoryController controller.CategoriesController) {
	r.POST("/category/add", categoryController.Create)
	r.PUT("/category/update/:categoryId", categoryController.Update)
	r.DELETE("/category/delete/:categoryId", categoryController.Delete)
	r.GET("/category/:categoryId", categoryController.FindById)
	r.GET("/category", categoryController.FindAll)
}

func NewBookRouter(r *gin.Engine, bookController controller.BookController) {
	r.POST("/book/add", bookController.Create)
	r.PUT("/book/update/:bookId", bookController.Update)
	r.DELETE("/book/delete/:bookId", bookController.Delete)
	r.GET("/book/:bookId", bookController.FindById)
	r.GET("/book", bookController.FindAll)
}

func NewOfficerRouter(r *gin.Engine, officerController controller.OfficerController) {
	r.POST("/officer/add", officerController.Create)
	r.PUT("/officer/update/:officerId", officerController.Update)
	r.DELETE("/officer/delete/:officerId", officerController.Delete)
	r.GET("/officer/:officerId", officerController.FindById)
	r.GET("/officer", officerController.FindAll)
}

func NewGuestBookRouter(r *gin.Engine, guestbookController controller.GuestBookController) {
	r.POST("/guestbook/add", guestbookController.Create)
	r.PUT("/guestbook/update/:guestbookId", guestbookController.Update)
	r.DELETE("/guestbook/delete/:guestbookId", guestbookController.Delete)
	r.GET("/guestbook/:guestbookId", guestbookController.FindById)
	r.GET("/guestbook", guestbookController.FindAll)
	r.GET("/guestbook/user/:userId", guestbookController.FindByUserId)
}
