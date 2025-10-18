// @title Auth Module API
// @version 1.0
// @description Authentication and authorization API
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api
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
	// Create User Repo
	userRepo := userRepo.NewUserRepository(db)
	// Create Auth Repo
	authRepo := repository.NewAuthRepository(db)
	// Create JWT Service
	jwtService := service.NewJWTService()
	// Create Auth Service and Handler
	authService := service.NewAuthService(authRepo, userRepo, jwtService)
	authHandler := handler.NewAuthHandler(authService)
	// Register Auth Routes
	routes.RegisterAuthRoutes(r, authHandler)
}

// AddJWTMiddleware adds the JWT middleware to the Gin engine
//
// Params:
//   - r: Gin engine instance
func AddJWTMiddleware(r *gin.Engine) {
	r.Use(middleware.JwtMiddleware())
}
