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
	auth.InitAuthModule(r, db)
	r.Run()
}
