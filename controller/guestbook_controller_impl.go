package controller

import (
	"github.com/dimassfeb-09/restapi-perpustakaan/exception"
	"github.com/dimassfeb-09/restapi-perpustakaan/helper"
	"github.com/dimassfeb-09/restapi-perpustakaan/model/web/guestbook"
	"github.com/dimassfeb-09/restapi-perpustakaan/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type GuestBookControllerImpl struct {
	GuestBookService service.GuestBookService
}

func NewGuestBookControllerImpl(guestBookService service.GuestBookService) GuestBookController {
	return &GuestBookControllerImpl{GuestBookService: guestBookService}
}

func (controller *GuestBookControllerImpl) Create(c *gin.Context) {

	xApiKey := c.Request.Header.Get("X-API-KEY")
	if xApiKey != "RAHASIA" {
		panic(exception.NewErrorUnauthorized("X-API-KEY Required."))
	}
	c.Writer.Header().Add("X-API-KEY", xApiKey)

	var createRequest guestbook.GuestBookCreateRequest
	err := c.ShouldBind(&createRequest)
	if err != nil {
		panic(exception.NewErrorShouldBind(err.Error()))
	}

	guestbookResponse := controller.GuestBookService.Create(c.Request.Context(), createRequest)

	webResponse := helper.WebResponse(http.StatusOK, "OK", guestbookResponse)
	c.JSON(http.StatusOK, webResponse)
}

func (controller *GuestBookControllerImpl) Update(c *gin.Context) {

	xApiKey := c.Request.Header.Get("X-API-KEY")
	if xApiKey != "RAHASIA" {
		panic(exception.NewErrorUnauthorized("X-API-KEY Required."))
	}
	c.Writer.Header().Add("X-API-KEY", xApiKey)

	var updateRequest guestbook.GuestBookUpdateRequest
	errMsg := c.ShouldBind(&updateRequest)
	if errMsg != nil {
		panic(exception.NewErrorShouldBind(errMsg.Error()))
	}

	guestbookId := c.Param("guestbookId")
	guestbookIdInt, err := strconv.Atoi(guestbookId)
	helper.PanicIfError(err)

	updateRequest.Id = guestbookIdInt
	guestbookResponse := controller.GuestBookService.Update(c.Request.Context(), updateRequest)

	webResponse := helper.WebResponse(http.StatusOK, "OK", guestbookResponse)
	c.JSON(http.StatusOK, webResponse)
}

func (controller *GuestBookControllerImpl) Delete(c *gin.Context) {

	xApiKey := c.Request.Header.Get("X-API-KEY")
	if xApiKey != "RAHASIA" {
		panic(exception.NewErrorUnauthorized("X-API-KEY Required."))
	}
	c.Writer.Header().Add("X-API-KEY", xApiKey)

	guestbookId := c.Param("guestbookId")
	guestbookIdInt, err := strconv.Atoi(guestbookId)
	helper.PanicIfError(err)

	findById := controller.GuestBookService.FindById(c.Request.Context(), guestbookIdInt)
	helper.PanicIfError(err)

	controller.GuestBookService.Delete(c.Request.Context(), findById.Id)

	webResponse := helper.WebResponse(http.StatusOK, "OK", "Berhasil hapus ID "+strconv.Itoa(guestbookIdInt))
	c.JSON(http.StatusOK, webResponse)

}

func (controller *GuestBookControllerImpl) FindById(c *gin.Context) {

	xApiKey := c.Request.Header.Get("X-API-KEY")
	if xApiKey != "RAHASIA" {
		panic(exception.NewErrorUnauthorized("X-API-KEY Required."))
	}
	c.Writer.Header().Add("X-API-KEY", xApiKey)

	guestbookId := c.Param("guestbookId")
	guestbookIdInt, err := strconv.Atoi(guestbookId)
	helper.PanicIfError(err)

	guestbookResponse := controller.GuestBookService.FindById(c.Request.Context(), guestbookIdInt)

	webResponse := helper.WebResponse(http.StatusOK, "OK", guestbookResponse)
	c.JSON(http.StatusOK, webResponse)
}

func (controller *GuestBookControllerImpl) FindAll(c *gin.Context) {

	xApiKey := c.Request.Header.Get("X-API-KEY")
	if xApiKey != "RAHASIA" {
		panic(exception.NewErrorUnauthorized("X-API-KEY Required."))
	}
	c.Writer.Header().Add("X-API-KEY", xApiKey)

	guestbookResponses := controller.GuestBookService.FindAll(c.Request.Context())
	webResponse := helper.WebResponse(http.StatusOK, "OK", guestbookResponses)
	c.JSON(http.StatusOK, webResponse)

}

func (controller *GuestBookControllerImpl) FindByUserId(c *gin.Context) {

	xApiKey := c.Request.Header.Get("X-API-KEY")
	if xApiKey != "RAHASIA" {
		panic(exception.NewErrorUnauthorized("X-API-KEY Required."))
	}
	c.Writer.Header().Add("X-API-KEY", xApiKey)

	userId := c.Param("userId")
	userIdInt, err := strconv.Atoi(userId)
	helper.PanicIfError(err)

	guestBookResponsebyUserIds := controller.GuestBookService.FindByUserId(c.Request.Context(), userIdInt)
	webResponse := helper.WebResponse(http.StatusOK, "OK", guestBookResponsebyUserIds)
	c.JSON(http.StatusOK, webResponse)

}
