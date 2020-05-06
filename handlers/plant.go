package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hydrophonics/controllers"
	"github.com/hydrophonics/types"
	"gopkg.in/birkirb/loggers.v1/log"
)

func AddPlants() gin.HandlerFunc {
	return func(c *gin.Context) {
		var plant types.PlantDTO

		err := c.ShouldBindJSON(&plant)
		if err != nil {
			c.AbortWithStatusJSON(400, "couldn't bind request data")
			return
		}
		response, err := controllers.AddPlants(plant)
		if err != nil {
			c.AbortWithStatusJSON(500, err)
			return
		}

		c.JSON(200, response)
	}
}

func GetPlants() gin.HandlerFunc {
	return func(c *gin.Context) {
		response, err := controllers.GetPlants()
		if err != nil {
			c.AbortWithStatusJSON(500, err)
			return
		}

		c.JSON(200, response)
	}
}

// func GetPlant() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		plantID := c.Param("plantID")
// 		pID, err := strconv.Atoi(plantID)
// 		if err != nil {
// 			log.Println("error strconv :", err)
// 			c.AbortWithStatusJSON(400, "invalid request")
// 			return
// 		}
// 		response, err := controllers.GetPlant()
// 		if err != nil {
// 			c.AbortWithStatusJSON(500, err)
// 			return
// 		}

// 		c.JSON(200, response)
// 	}
// }

func DeletePlant() gin.HandlerFunc {
	return func(c *gin.Context) {
		plantID := c.Param("plantID")
		pID, err := strconv.Atoi(plantID)
		if err != nil {
			log.Println("error strconv :", err)
			c.AbortWithStatusJSON(400, "invalid request")
			return
		}

		response, err := controllers.DeletePlant(uint(pID))
		if err != nil {
			c.AbortWithStatusJSON(500, err)
			return
		}

		c.JSON(200, response)
	}
}
