package routes

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// RegisterAuthSwaggerRoutes registers the Swagger documentation route for the auth module
// This allows the auth module to have its own Swagger documentation endpoint
//
// Usage:
//
//	routes.RegisterAuthSwaggerRoutes(r)
//
// This will make the auth module docs available at: /auth-docs/swagger/index.html
func RegisterAuthSwaggerRoutes(r *gin.Engine) {
	authDocsGroup := r.Group("/auth-docs")
	{
		// Serve Swagger UI
		authDocsGroup.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
}
