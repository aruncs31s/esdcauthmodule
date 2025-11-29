package auth

import (
	azf "github.com/aruncs31s/azf"
	azfRoutes "github.com/aruncs31s/azf/application/routes"
	"github.com/aruncs31s/esdcauthmodule/handler"
	"github.com/aruncs31s/esdcauthmodule/middleware"
	"github.com/aruncs31s/esdcauthmodule/repository"
	"github.com/aruncs31s/esdcauthmodule/routes"
	"github.com/aruncs31s/esdcauthmodule/service"
	userRepo "github.com/aruncs31s/esdcusermodule/repository"
	"github.com/casbin/casbin/v2"
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

// InitAZFModule initializes the AZF Authorization Framework
// This sets up Casbin-based RBAC authorization
//
// Params:
//   - db: GORM database instance (optional, can be nil)
//   - enforcer: Casbin enforcer instance (optional, can be nil)
func InitAZFModule(db *gorm.DB, enforcer *casbin.Enforcer) {
	azf.InitAuthZModule(db, enforcer)
}

// InitAZFUsageTracking initializes API usage tracking
// Call this after InitAZFModule
func InitAZFUsageTracking() {
	azf.InitUsageTracking()
}

// AddAZFTrackingMiddleware adds API usage tracking middleware
// Tracks endpoint usage, response times, and other metrics
//
// Params:
//   - r: Gin engine instance
func AddAZFTrackingMiddleware(r *gin.Engine) {
	azf.SetApiTrackingMiddleware(r)
}

// AddAZFAuthMiddleware adds AZF authorization middleware
// Enforces Casbin-based RBAC policies
//
// Params:
//   - r: Gin engine instance
func AddAZFAuthMiddleware(r *gin.Engine) {
	azf.SetAuthZMiddleware(r)
}

// SetupAZFUI sets up the AZF admin UI
// Provides web interface for policy management and analytics
//
// Params:
//   - r: Gin engine instance
func SetupAZFUI(r *gin.Engine) {
	azf.SetupUI(r)
}

// SetupAZFDocs sets up documentation routes
//
// Params:
//   - r: Gin engine instance
//   - docsPath: Path to the documentation directory
func SetupAZFDocs(r *gin.Engine, docsPath string) {
	azfRoutes.SetupDocsRoutes(r, docsPath)
}

// InitFullAZF is a convenience function that initializes all AZF components
// This includes: authorization module, usage tracking, tracking middleware,
// authorization middleware, and admin UI
//
// Params:
//   - r: Gin engine instance
//   - db: GORM database instance (optional, can be nil)
//   - enforcer: Casbin enforcer instance (optional, can be nil)
func InitFullAZF(r *gin.Engine, db *gorm.DB, enforcer *casbin.Enforcer) {
	// Initialize AZF Authorization Module
	InitAZFModule(db, enforcer)

	// Initialize Usage Tracking
	InitAZFUsageTracking()

	// Add API Tracking Middleware
	AddAZFTrackingMiddleware(r)

	// Setup Admin UI
	SetupAZFUI(r)

	// Add Authorization Middleware
	AddAZFAuthMiddleware(r)
}
