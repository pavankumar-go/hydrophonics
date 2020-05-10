package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hydrophonics/controllers"
	"github.com/hydrophonics/types"
	"gopkg.in/birkirb/loggers.v1/log"
)

// RegisterGrowthInfoRoutes defines growth-info routes
func RegisterGrowthInfoRoutes(api *gin.RouterGroup) {
	api.POST("/growthinfo/", addGrowthInfo())
	api.GET("/plant/:plantID/info/", getGrowthInfoForPlant())
	api.DELETE("/growthinfo/:growthinfoID/plant/:plantID/", deleteGrowthInfoOfPlant())
}

// AddGrowthInfo adds info to a plant
func addGrowthInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		var info types.GrowthDTO

		err := c.ShouldBindJSON(&info)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "invalid request data"})
			return
		}

		response, err := controllers.AddGrowthInfo(info)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, response)
	}
}

// GetGrowthInfoForPlant lists the growth info of a plant
func getGrowthInfoForPlant() gin.HandlerFunc {
	return func(c *gin.Context) {
		plantID := c.Param("plantID")
		pID, err := strconv.Atoi(plantID)
		if err != nil {
			log.Println("error strconv :", err)
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
			return
		}

		response, err := controllers.GetGrowthInfoForPlant(uint(pID))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, response)
	}
}

// GetGrowthInfoForPlant lists the growth info of a plant
func deleteGrowthInfoOfPlant() gin.HandlerFunc {
	return func(c *gin.Context) {
		plantID := c.Param("plantID")
		growthInfoID := c.Param("growthinfoID")

		pID, err := strconv.Atoi(plantID)
		if err != nil {
			log.Println("error strconv :", err)
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
			return
		}

		gID, err := strconv.Atoi(growthInfoID)
		if err != nil {
			log.Println("error strconv :", err)
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
			return
		}

		response, err := controllers.DeleteGrowthInfoOfPlant(uint(pID), uint(gID))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, response)
	}
}
