package helper

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func MiddleWare(c *gin.Context) {
	apiKey := c.Request.Header.Get("X-API-KEY")
	if apiKey != "RAHASIA" {
		c.JSON(http.StatusUnauthorized, WebResponse(http.StatusUnauthorized, "Unauthorized", gin.H{"error": "X-API-KEY are Required"}))
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	c.Writer.Header().Add("X-API-KEY", apiKey)
}
