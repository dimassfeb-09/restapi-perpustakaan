package app

import (
	"github.com/dimassfeb-09/restapi-perpustakaan/controller"
	"github.com/gin-gonic/gin"
)

func NewUserRouter(r *gin.Engine, userController controller.UserController) {
	r.POST("/user/add", userController.Create)
	r.PUT("/user/update/:id", userController.Update)
	r.DELETE("/user/delete/:id", userController.Delete)
	r.GET("/user/:id", userController.FindById)
	r.GET("/user/filter/:filterBy/:value", userController.FindBy)
}
