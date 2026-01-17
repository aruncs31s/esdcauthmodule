package handler

import "github.com/gin-gonic/gin"

func ShowLoginPage(c *gin.Context) {
	c.Header("Content-Type", "text/html")
	c.File("./static/common/login.html")
}

func ShowRegisterPage(c *gin.Context) {
	c.Header("Content-Type", "text/html")
	c.File("./static/common/register.html")
}
