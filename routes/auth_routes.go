package routes

import (
	"github.com/aruncs31s/esdcauthmodule/handler"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(r *gin.Engine, authHandler handler.AuthHandler) {
	userRoutes := r.Group("/api/user")
	{
		userRoutes.POST("/login", authHandler.Login)
		userRoutes.POST("/register", authHandler.Register)
	}
	authPages := r.Group("/auth")
	{
		authPages.GET("/login", handler.ShowLoginPage)
		authPages.GET("/register", handler.ShowRegisterPage)
	}
}
