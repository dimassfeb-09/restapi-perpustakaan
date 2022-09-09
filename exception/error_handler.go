package exception

import (
	"github.com/dimassfeb-09/restapi-perpustakaan/model/web"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ErrorHandler(c *gin.Context, recovered interface{}) {

	if errorDuplicate(c, recovered) == true {
		return
	}

	if notFoundErr(c, recovered) == true {
		return
	}

	if internalServerError(c, recovered) == true {
		return
	}

}

func notFoundErr(c *gin.Context, recovered interface{}) bool {
	err, ok := recovered.(NotFoundError)

	if ok {
		webResponse := web.WebResponse{
			Code:   http.StatusNotFound,
			Status: "Bad Not Found",
			Data:   err,
		}
		c.JSON(http.StatusNotFound, webResponse)
		return true
	}

	return false
}

func errorDuplicate(c *gin.Context, recovered interface{}) bool {
	err, ok := recovered.(ErrorDuplicate)

	if ok {
		webResponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   err,
		}
		c.JSON(http.StatusBadRequest, webResponse)
		return true
	}

	return false
}

func internalServerError(c *gin.Context, recovered interface{}) bool {
	err, ok := recovered.(string)

	if ok {
		webResponse := web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "Internal Server Error",
			Data:   err,
		}
		c.JSON(http.StatusInternalServerError, webResponse)
		return true
	}

	return false
}
