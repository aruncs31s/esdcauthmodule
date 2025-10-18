package auth

import (
	"github.com/aruncs31s/esdcauthmodule/handler"
	"github.com/aruncs31s/esdcauthmodule/repository"
	"github.com/aruncs31s/esdcauthmodule/routes"
	"github.com/aruncs31s/esdcauthmodule/service"
	userRepo "github.com/aruncs31s/esdcusermodule/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitAuthModule(r *gin.Engine, db *gorm.DB) {
	userRepo := userRepo.NewUserRepository(db)
	authRepo := repository.NewAuthRepository(db)
	jwtService := service.NewJWTService()
	authService := service.NewAuthService(authRepo, userRepo, jwtService)
	authHandler := handler.NewAuthHandler(authService)
	routes.RegisterAuthRoutes(r, authHandler)
}
