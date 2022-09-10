package controller

import (
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

	err := c.ShouldBindJSON(&createRequest)
	helper.ErrorShouldBind(c, err, http.StatusBadRequest, "Status Bad Request")

	userResponse := controller.UserService.Create(c.Request.Context(), createRequest)

	webResponse := helper.WebResponse(http.StatusOK, "OK", userResponse)
	c.JSON(http.StatusOK, webResponse)
}

func (controller *UserControllerImpl) Update(c *gin.Context) {

	var updateRequest user.UserUpdateRequest
	err := c.ShouldBindJSON(&updateRequest)
	helper.ErrorShouldBind(c, err, http.StatusBadRequest, "Status Bad Request")

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
