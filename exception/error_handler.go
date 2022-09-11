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

	if errorBadRequest(c, recovered) == true {
		return
	}

	if errorForbidden(c, recovered) == true {
		return
	}

	if notFoundErr(c, recovered) == true {
		return
	}

	if internalServerError(c, recovered) == true {
		return
	}

	if errorShouldBind(c, recovered) == true {
		return
	}

	if errorUnauthorized(c, recovered) == true {
		return
	}

}

func errorForbidden(c *gin.Context, recovered interface{}) bool {
	err, ok := recovered.(ErrorForbidden)

	if ok {
		webResponse := helper.WebResponse(http.StatusForbidden, "Forbidden", err)
		c.JSON(http.StatusForbidden, webResponse)
		return true
	}
	return false
}

func errorUnauthorized(c *gin.Context, recovered interface{}) bool {
	err, ok := recovered.(ErrorUnauthorized)

	if ok {
		webResponse := helper.WebResponse(http.StatusUnauthorized, "Unauthorized", err)
		c.JSON(http.StatusUnauthorized, webResponse)
		return true
	}
	return false
}

func errorShouldBind(c *gin.Context, recovered interface{}) bool {
	err, ok := recovered.(ErrorShouldBind)

	if ok {
		webResponse := helper.WebResponse(http.StatusBadRequest, "Bad Request", err)
		c.JSON(http.StatusBadRequest, webResponse)
		return true
	}
	return false
}

func errorBadRequest(c *gin.Context, recovered interface{}) bool {
	err, ok := recovered.(ErrorBadRequest)

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
