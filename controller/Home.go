package controller

import "github.com/gin-gonic/gin"

func HandleHome(c *gin.Context) {
	c.JSON(200, "Home Page")
}
