package controller

import (
	"github.com/dimassfeb-09/restapi-perpustakaan/exception"
	"github.com/dimassfeb-09/restapi-perpustakaan/helper"
	"github.com/dimassfeb-09/restapi-perpustakaan/model/web/user"
	"github.com/dimassfeb-09/restapi-perpustakaan/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserControllerImpl struct {
	UserService service.UserService
}

func NewUserControllerImpl(userService service.UserService) UserController {
	return &UserControllerImpl{UserService: userService}
}

func (controller *UserControllerImpl) Create(c *gin.Context) {

	var createRequest user.UserCreateRequest
	err := c.ShouldBind(&createRequest)
	if err != nil {
		panic(exception.NewErrorShouldBind(err.Error()))
	}

	userResponse := controller.UserService.Create(c.Request.Context(), createRequest)

	webResponse := helper.WebResponse(http.StatusOK, "OK", userResponse)
	c.JSON(http.StatusOK, webResponse)
}

func (controller *UserControllerImpl) Update(c *gin.Context) {

	var updateRequest user.UserUpdateRequest
	err := c.ShouldBind(&updateRequest)
	if err != nil {
		panic(exception.NewErrorShouldBind(err.Error()))
	}

	userId := c.Param("id")
	userIdInt, err := strconv.Atoi(userId)
	helper.PanicIfError(err)

	updateRequest.Id = userIdInt
	userResponse := controller.UserService.Update(c.Request.Context(), updateRequest)

	webResponse := helper.WebResponse(http.StatusOK, "OK", userResponse)
	c.JSON(http.StatusOK, webResponse)
}

func (controller *UserControllerImpl) Delete(c *gin.Context) {

	userId := c.Param("id")
	userIdInt, err := strconv.Atoi(userId)
	helper.PanicIfError(err)

	findById, err := controller.UserService.FindById(c.Request.Context(), userIdInt)
	helper.PanicIfError(err)

	controller.UserService.Delete(c.Request.Context(), findById.Id)

	webResponse := helper.WebResponse(http.StatusOK, "OK", "Berhasil hapus ID "+strconv.Itoa(userIdInt))
	c.JSON(http.StatusOK, webResponse)

}

func (controller *UserControllerImpl) FindById(c *gin.Context) {

	userId := c.Param("id")
	userIdInt, err := strconv.Atoi(userId)
	helper.PanicIfError(err)

	userResponse, err := controller.UserService.FindById(c.Request.Context(), userIdInt)
	helper.PanicIfError(err)

	webResponse := helper.WebResponse(http.StatusOK, "OK", userResponse)
	c.JSON(http.StatusOK, webResponse)
}

func (controller *UserControllerImpl) FindAll(c *gin.Context) {

	userResponses := controller.UserService.FindAll(c.Request.Context())
	webResponse := helper.WebResponse(http.StatusOK, "OK", userResponses)
	c.JSON(http.StatusOK, webResponse)

}

func (controller *UserControllerImpl) LoginAuth(c *gin.Context) {

	var loginAuth user.LoginAuth
	err := c.ShouldBind(&loginAuth)
	if err != nil {
		panic(exception.NewErrorShouldBind(err.Error()))
	}

	userName := c.PostForm("username")
	passWord := c.PostForm("password")

	userResponse, err := controller.UserService.LoginAuth(c.Request.Context(), userName, passWord)
	if err != nil {
		panic(exception.NewErrorBadRequest(err.Error()))
	}

	webResponse := helper.WebResponse(http.StatusOK, "OK", userResponse)
	c.JSON(http.StatusOK, webResponse)
}
