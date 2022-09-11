package controller

import (
	"github.com/dimassfeb-09/restapi-perpustakaan/exception"
	"github.com/dimassfeb-09/restapi-perpustakaan/helper"
	"github.com/dimassfeb-09/restapi-perpustakaan/model/web/officer"
	"github.com/dimassfeb-09/restapi-perpustakaan/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type OfficerControllerImpl struct {
	OfficerService service.OfficerService
}

func NewOfficerControllerImpl(officerService service.OfficerService) OfficerController {
	return &OfficerControllerImpl{OfficerService: officerService}
}

func (controller *OfficerControllerImpl) Create(c *gin.Context) {
	var officerRequestCreate officer.OfficerCreateRequest

	err := c.ShouldBindJSON(&officerRequestCreate)
	if err != nil {
		panic(exception.NewErrorShouldBind(err.Error()))
	}

	officerResponse := controller.OfficerService.Create(c.Request.Context(), officerRequestCreate)

	webResponse := helper.WebResponse(http.StatusOK, "OK", officerResponse)
	c.JSON(http.StatusOK, webResponse)
}

func (controller *OfficerControllerImpl) Update(c *gin.Context) {
	officerId := c.Param("officerId")
	officerIdInt, err := strconv.Atoi(officerId)
	helper.PanicIfError(err)

	var officerUpdateRequest officer.OfficerUpdateRequest

	errMsg := c.ShouldBindJSON(&officerUpdateRequest)
	if errMsg != nil {
		panic(exception.NewErrorShouldBind(errMsg.Error()))
	}

	officerUpdateRequest.Id = officerIdInt
	officerResponse := controller.OfficerService.Update(c.Request.Context(), officerUpdateRequest)
	webResponse := helper.WebResponse(http.StatusOK, "OK", officerResponse)
	c.JSON(http.StatusOK, webResponse)
}

func (controller *OfficerControllerImpl) Delete(c *gin.Context) {
	officerId := c.Param("officerId")
	officerIdInt, err := strconv.Atoi(officerId)
	helper.PanicIfError(err)

	controller.OfficerService.Delete(c.Request.Context(), officerIdInt)

	webResponse := helper.WebResponse(http.StatusOK, "OK", "Office dengan ID "+officerId+" berhasil dihapus")
	c.JSON(http.StatusOK, webResponse)
}

func (controller *OfficerControllerImpl) FindById(c *gin.Context) {

	officerId := c.Param("officerId")
	officerIdInt, err := strconv.Atoi(officerId)
	helper.PanicIfError(err)

	officerResponse := controller.OfficerService.FindById(c.Request.Context(), officerIdInt)

	webResponse := helper.WebResponse(http.StatusOK, "OK", officerResponse)
	c.JSON(http.StatusOK, webResponse)
}

func (controller *OfficerControllerImpl) FindAll(c *gin.Context) {
	officerResponses := controller.OfficerService.FindAll(c.Request.Context())

	webResponse := helper.WebResponse(http.StatusOK, "OK", officerResponses)
	c.JSON(http.StatusOK, webResponse)
}
