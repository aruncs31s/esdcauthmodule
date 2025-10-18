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
