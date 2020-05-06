package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hydrophonics/controllers"
	"github.com/hydrophonics/types"
	"gopkg.in/birkirb/loggers.v1/log"
)

func AddGrowthInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		var info types.GrowthDTO

		err := c.ShouldBindJSON(&info)
		if err != nil {
			c.AbortWithStatusJSON(400, "couldn't bind request data")
			return
		}
		response, err := controllers.AddGrowthInfo(info)
		if err != nil {
			c.AbortWithStatusJSON(500, err)
			return
		}

		c.JSON(200, response)
	}
}

func GetGrowthInfoForPlant() gin.HandlerFunc {
	return func(c *gin.Context) {
		plantID := c.Param("plantID")
		pID, err := strconv.Atoi(plantID)
		if err != nil {
			log.Println("error strconv :", err)
			c.AbortWithStatusJSON(400, "invalid request")
			return
		}
		response, err := controllers.GetGrowthInfoForPlant(uint(pID))
		if err != nil {
			c.AbortWithStatusJSON(500, err)
			return
		}

		c.JSON(200, response)
	}
}
