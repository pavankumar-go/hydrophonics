package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hydrophonics/controllers"
)

// RegisterCategoryRoutes registers category routes
func RegisterCategoryRoutes(api *gin.RouterGroup) {
	api.GET("/categories/", getCategories())
}

// GetCategories lists all plant categories
func getCategories() gin.HandlerFunc {
	return func(c *gin.Context) {
		categories, err := controllers.GetCategories()
		if err != nil {
			log.Println("getCategories : ", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "could not get categories",
			})
			return
		}

		c.JSON(http.StatusOK, categories)
	}
}
