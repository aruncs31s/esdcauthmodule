package auth

import (
	"github.com/aruncs31s/esdcauthmodule/handler"
	"github.com/aruncs31s/esdcauthmodule/middleware"
	"github.com/aruncs31s/esdcauthmodule/repository"
	"github.com/aruncs31s/esdcauthmodule/routes"
	"github.com/aruncs31s/esdcauthmodule/service"
	userRepo "github.com/aruncs31s/esdcusermodule/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// InitAuthModule initializes the authentication module
//
// Params:
//   - r: Gin engine instance
//   - db: GORM database instance
func InitAuthModule(r *gin.Engine, db *gorm.DB) {
	userRepo := userRepo.NewUserRepository(db)
	authRepo := repository.NewAuthRepository(db)
	jwtService := service.NewJWTService()
	authService := service.NewAuthService(authRepo, userRepo, jwtService)
	authHandler := handler.NewAuthHandler(authService)
	routes.RegisterAuthRoutes(r, authHandler)
}

// AddJWTMiddleware adds the JWT middleware to the Gin engine
//
// Params:
//   - r: Gin engine instance
func AddJWTMiddleware(r *gin.Engine) {
	r.Use(middleware.JwtMiddleware())
}

// SetupSwagger sets up the Swagger documentation routes for the auth module
//
// Call this function in your main application after initializing the auth module
// to enable Swagger UI at /swagger/index.html
//
// Params:
//   - r: Gin engine instance
//
// Example:
//
//	auth.InitAuthModule(r, db)
//	auth.SetupSwagger(r)
//	auth.AddJWTMiddleware(r)
func SetupSwagger(r *gin.Engine) {
	// This function prepares the module for Swagger documentation
	// The actual Swagger setup should be done in the consuming application
	// using swag or a similar tool to generate swagger docs
}
