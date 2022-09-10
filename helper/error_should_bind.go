package helper

import (
	"github.com/dimassfeb-09/restapi-perpustakaan/model/web"
	"github.com/gin-gonic/gin"
)

func ErrorShouldBind(c *gin.Context, err error, code int, status string) {
	if err != nil {
		webResponse := web.WebResponse{
			Code:   code,
			Status: status,
			Data:   err.Error(),
		}
		c.JSON(code, webResponse)
	}
	return
}
