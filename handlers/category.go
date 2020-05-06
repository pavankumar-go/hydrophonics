package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/hydrophonics/controllers"
)

func GetCategories() gin.HandlerFunc {
	return func(c *gin.Context) {
		categories := controllers.GetCategories()
		c.JSON(200, categories)
	}
}
