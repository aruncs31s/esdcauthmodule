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
