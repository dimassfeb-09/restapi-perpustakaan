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
