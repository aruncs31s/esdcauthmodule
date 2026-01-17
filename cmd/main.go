package main

import (
	"log"

	auth "github.com/aruncs31s/esdcauthmodule"
	model "github.com/aruncs31s/esdcmodels"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	r := gin.Default()
	db, err := gorm.Open(sqlite.Open("db.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	if err := db.AutoMigrate(model.User{}); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
	if err := db.AutoMigrate(model.Github{}); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// Initialize Auth Module (login/register routes)
	auth.InitAuthModule(r, db)

	// Set user context middleware (customize based on your auth flow)
	r.Use(func(c *gin.Context) {
		// This would typically come from JWT claims after authentication
		// For testing, we set static values
		if claims, exists := c.Get("jwt_claims"); exists {
			if claimsMap, ok := claims.(map[string]interface{}); ok {
				if userID, ok := claimsMap["user_id"].(string); ok {
					c.Set("user_id", userID)
				}
				if role, ok := claimsMap["role"].(string); ok {
					c.Set("role", role)
				}
			}
		}
		c.Next()
	})

	// // Initialize AZF Authorization Framework with all features
	// // This includes: authorization, usage tracking, admin UI
	// auth.InitFullAZF(r, db, nil) // Pass nil for Casbin enforcer to use default
	//
	// Protected routes go here
	r.GET("/api/protected", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "protected resource"})
	})

	r.Run()
}
