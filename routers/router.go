package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/hydrophonics/handlers"
)

// StartServer registers routes and starts server
func StartServer() {

	app := gin.New()
	api := app.Group("/api/v1")

	handlers.RegisterPlantRoutes(api)
	handlers.RegisterGrowthInfoRoutes(api)
	handlers.RegisterCategoryRoutes(api)

	api.Use(gin.Logger())
	app.Run()
}
