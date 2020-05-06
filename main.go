package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hydrophonics/database"
	"github.com/hydrophonics/database/migrate"
	"github.com/hydrophonics/database/seed"
	"github.com/hydrophonics/handlers"
)

func main() {

	db := database.Init()
	defer db.Close()
	migrate.Migrate(db)
	seed.Category(db)

	app := gin.New()
	api := app.Group("/api/v1")
	api.POST("/plants/", handlers.AddPlants())
	api.GET("/plants/", handlers.GetPlants())
	api.GET("/plants/:plantID/info/", handlers.GetGrowthInfoForPlant())
	api.DELETE("/plants/:plantID/", handlers.DeletePlant())

	api.POST("/info/", handlers.AddGrowthInfo())

	api.GET("/categories/", handlers.GetCategories())
	api.Use(gin.Logger())
	app.Run(":8080")

}
