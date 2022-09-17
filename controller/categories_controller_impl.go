package controller

import (
	"github.com/dimassfeb-09/restapi-perpustakaan/exception"
	"github.com/dimassfeb-09/restapi-perpustakaan/helper"
	"github.com/dimassfeb-09/restapi-perpustakaan/model/web/categories"
	"github.com/dimassfeb-09/restapi-perpustakaan/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CategoriesControllerImpl struct {
	CategoriesService service.CategoriesService
}

func NewCategoriesControllerImpl(categoriesService service.CategoriesService) CategoriesController {
	return &CategoriesControllerImpl{CategoriesService: categoriesService}
}

func (controller *CategoriesControllerImpl) Create(c *gin.Context) {

	var categoriesRequestCreate categories.CategoriesCreateRequest
	err := c.ShouldBind(&categoriesRequestCreate)
	if err != nil {
		panic(exception.NewErrorShouldBind(err.Error()))
	}

	categoriesResponse := controller.CategoriesService.Create(c.Request.Context(), categoriesRequestCreate)

	webResponse := helper.WebResponse(http.StatusOK, "OK", categoriesResponse)
	c.JSON(http.StatusOK, webResponse)
}

func (controller *CategoriesControllerImpl) Update(c *gin.Context) {

	categoryId := c.Param("categoryId")
	categoryIdInt, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	var categoriesUpdateRequest categories.CategoriesUpdateRequest

	errMsg := c.ShouldBind(&categoriesUpdateRequest)
	if errMsg != nil {
		panic(exception.NewErrorShouldBind(errMsg.Error()))
	}

	categoriesUpdateRequest.Id = categoryIdInt
	categoriesResponse := controller.CategoriesService.Update(c.Request.Context(), categoriesUpdateRequest)
	webResponse := helper.WebResponse(http.StatusOK, "OK", categoriesResponse)
	c.JSON(http.StatusOK, webResponse)
}

func (controller *CategoriesControllerImpl) Delete(c *gin.Context) {

	categoryId := c.Param("categoryId")
	categoryIdInt, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	controller.CategoriesService.Delete(c.Request.Context(), categoryIdInt)

	webResponse := helper.WebResponse(http.StatusOK, "OK", "Category ID dengan "+categoryId+" berhasil dihapus")
	c.JSON(http.StatusOK, webResponse)
}

func (controller *CategoriesControllerImpl) FindById(c *gin.Context) {

	categoryId := c.Param("categoryId")
	categoryIdInt, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	categoriesResponse := controller.CategoriesService.FindById(c.Request.Context(), categoryIdInt)

	webResponse := helper.WebResponse(http.StatusOK, "OK", categoriesResponse)
	c.JSON(http.StatusOK, webResponse)
}

func (controller *CategoriesControllerImpl) FindAll(c *gin.Context) {

	categoriesResponses := controller.CategoriesService.FindAll(c.Request.Context())
	webResponse := helper.WebResponse(http.StatusOK, "OK", categoriesResponses)
	c.JSON(http.StatusOK, webResponse)
}
