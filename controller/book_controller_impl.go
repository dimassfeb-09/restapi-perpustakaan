package controller

import (
	"github.com/dimassfeb-09/restapi-perpustakaan/exception"
	"github.com/dimassfeb-09/restapi-perpustakaan/helper"
	"github.com/dimassfeb-09/restapi-perpustakaan/model/web/book"
	"github.com/dimassfeb-09/restapi-perpustakaan/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type BookControllerImpl struct {
	BookService service.BookService
}

func NewBookControllerImpl(bookService service.BookService) BookController {
	return &BookControllerImpl{BookService: bookService}
}

func (controller *BookControllerImpl) Create(c *gin.Context) {
	var bookRequestCreate book.BookCreateRequest

	err := c.ShouldBindJSON(&bookRequestCreate)
	if err != nil {
		panic(exception.NewErrorShouldBind(err.Error()))
	}

	bookResponse := controller.BookService.Create(c.Request.Context(), bookRequestCreate)

	webResponse := helper.WebResponse(http.StatusOK, "OK", bookResponse)
	c.JSON(http.StatusOK, webResponse)
}

func (controller *BookControllerImpl) Update(c *gin.Context) {
	bookId := c.Param("bookId")
	bookIdInt, err := strconv.Atoi(bookId)
	helper.PanicIfError(err)

	var bookUpdateRequest book.BookUpdateRequest

	errMsg := c.ShouldBindJSON(&bookUpdateRequest)
	if errMsg != nil {
		panic(exception.NewErrorShouldBind(errMsg.Error()))
	}

	bookUpdateRequest.Id = bookIdInt
	bookResponse := controller.BookService.Update(c.Request.Context(), bookUpdateRequest)
	webResponse := helper.WebResponse(http.StatusOK, "OK", bookResponse)
	c.JSON(http.StatusOK, webResponse)
}

func (controller *BookControllerImpl) Delete(c *gin.Context) {
	bookId := c.Param("bookId")
	bookIdInt, err := strconv.Atoi(bookId)
	helper.PanicIfError(err)

	controller.BookService.Delete(c.Request.Context(), bookIdInt)

	webResponse := helper.WebResponse(http.StatusOK, "OK", "Buku dengan ID "+bookId+" berhasil dihapus")
	c.JSON(http.StatusOK, webResponse)
}

func (controller *BookControllerImpl) FindById(c *gin.Context) {

	bookId := c.Param("bookId")
	bookIdInt, err := strconv.Atoi(bookId)
	helper.PanicIfError(err)

	bookResponse := controller.BookService.FindById(c.Request.Context(), bookIdInt)

	webResponse := helper.WebResponse(http.StatusOK, "OK", bookResponse)
	c.JSON(http.StatusOK, webResponse)
}

func (controller *BookControllerImpl) FindAll(c *gin.Context) {
	bookResponses := controller.BookService.FindAll(c.Request.Context())

	webResponse := helper.WebResponse(http.StatusOK, "OK", bookResponses)
	c.JSON(http.StatusOK, webResponse)
}
