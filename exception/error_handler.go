package exception

import (
	"github.com/dimassfeb-09/restapi-perpustakaan/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ErrorHandler(c *gin.Context, recovered interface{}) {

	if errorDataRegistered(c, recovered) == true {
		return
	}

	if errorInvalidDataType(c, recovered) == true {
		return
	}

	if notFoundErr(c, recovered) == true {
		return
	}

	if internalServerError(c, recovered) == true {
		return
	}

}

func errorInvalidDataType(c *gin.Context, recovered interface{}) bool {
	err, ok := recovered.(ErrorInvalidDataType)

	if ok {
		webResponse := helper.WebResponse(http.StatusBadRequest, "Bad Request", err)
		c.JSON(http.StatusBadRequest, webResponse)
		return true
	}

	return false
}

func notFoundErr(c *gin.Context, recovered interface{}) bool {
	err, ok := recovered.(NotFoundError)

	if ok {
		webResponse := helper.WebResponse(http.StatusNotFound, "Bad Not Found", err)
		c.JSON(http.StatusNotFound, webResponse)
		return true
	}

	return false
}

func errorDataRegistered(c *gin.Context, recovered interface{}) bool {
	err, ok := recovered.(ErrorDataRegistered)

	if ok {
		webResponse := helper.WebResponse(http.StatusBadRequest, "Bad Request", err)
		c.JSON(http.StatusBadRequest, webResponse)
		return true
	}

	return false
}

func internalServerError(c *gin.Context, recovered interface{}) bool {
	err, ok := recovered.(string)

	if ok {
		webResponse := helper.WebResponse(http.StatusInternalServerError, "Internal Server Error", err)
		c.JSON(http.StatusInternalServerError, webResponse)
		return true
	}

	return false
}
