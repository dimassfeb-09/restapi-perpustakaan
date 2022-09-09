package controller

import (
	"github.com/dimassfeb-09/restapi-perpustakaan/helper"
	"github.com/dimassfeb-09/restapi-perpustakaan/model/web"
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

	apiKey := c.Request.Header.Get("X-API-KEY")
	if apiKey == "" {
		helper.ErrorShouldBind(c, err, http.StatusUnauthorized, "Unauthorized X-API-KEY")
	}

	userResponse := controller.UserService.Create(c.Request.Context(), createRequest)

	c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   userResponse,
	})
}

func (controller *UserControllerImpl) Update(c *gin.Context) {

	var updateRequest user.UserUpdateRequest
	err := c.ShouldBindJSON(&updateRequest)
	helper.ErrorShouldBind(c, err, http.StatusBadRequest, "Status Bad Request")

	userId := c.Param("id")
	userIdInt, err := strconv.Atoi(userId)
	helper.PanicIfError(err)

	controller.UserService.FindBy(c.Request.Context(), "id", userIdInt)

	userResponse := controller.UserService.Update(c.Request.Context(), updateRequest)

	c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   userResponse,
	})
}

func (controller *UserControllerImpl) Delete(c *gin.Context) {

	userId := c.Param("id")
	userIdInt, err := strconv.Atoi(userId)
	helper.PanicIfError(err)

	findById, err := controller.UserService.FindById(c.Request.Context(), userIdInt)
	helper.PanicIfError(err)

	controller.UserService.Delete(c.Request.Context(), findById.Id)

	c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   "Berhasil hapus ID " + strconv.Itoa(userIdInt),
	})

}

func (controller *UserControllerImpl) FindBy(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (controller *UserControllerImpl) FindById(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}
