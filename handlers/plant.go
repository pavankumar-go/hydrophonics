package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hydrophonics/controllers"
	"github.com/hydrophonics/types"
	"gopkg.in/birkirb/loggers.v1/log"
)

// RegisterPlantRoutes defines routes for plant ops
func RegisterPlantRoutes(api *gin.RouterGroup) {
	api.POST("/plant/", addPlant())
	api.GET("/plants/", getPlants())
	api.GET("/plant/", getPlant())
	api.DELETE("/plant/:plantID/", deletePlant())
	api.PATCH("/plant/:plantID/", updatePlant())
}

// AddPlant adds a plant
func addPlant() gin.HandlerFunc {
	return func(c *gin.Context) {
		var plant types.PlantDTO

		err := c.ShouldBindJSON(&plant)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "request body is invalid",
			})
			return
		}

		response, err := controllers.AddPlant(plant)
		if err != nil {
			log.Println("addPlant : ", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "could not add plant",
			})
			return
		}

		c.JSON(http.StatusOK, response)
	}
}

// getPlants lists all plants
func getPlants() gin.HandlerFunc {
	return func(c *gin.Context) {
		response, err := controllers.GetPlants()
		if err != nil {
			log.Println("getPlants : ", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "could not get plants",
			})
			return
		}

		c.JSON(http.StatusOK, response)
	}
}

// getPlant returns a plant that matches the name
func getPlant() gin.HandlerFunc {
	return func(c *gin.Context) {
		plantName := c.Request.URL.Query().Get("name")
		if plantName == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "request is invalid",
			})
			return
		}

		response, err := controllers.GetPlant(plantName)
		if err != nil {
			log.Println("getPlant : ", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "could not get plant",
			})
			return
		}

		c.JSON(200, response)
	}
}

// deletePlant deletes a plant
func deletePlant() gin.HandlerFunc {
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
			if err.Error() == "record not found" {
				c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
					"message": "plant with ID: " + plantID + " not found",
				})
				return
			}
			log.Println("deletePlant : ", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "could not delete plant",
			})
			return
		}

		c.JSON(http.StatusOK, response)
	}
}

// updatePlant updates a plant
func updatePlant() gin.HandlerFunc {
	return func(c *gin.Context) {
		plantID := c.Param("plantID")
		pID, err := strconv.Atoi(plantID)
		if err != nil {
			log.Println("error strconv :", err)
			c.AbortWithStatusJSON(400, "invalid request")
			return
		}

		requestBody, err := ioutil.ReadAll(c.Request.Body)
		if err != nil || len(requestBody) <= 0 {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "request body is invalid",
			})
			return
		}

		var updateData types.UpdatePlantDTO
		err = json.Unmarshal(requestBody, &updateData)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "request body is invalid",
			})
			return
		}
		response, err := controllers.UpdatePlant(uint(pID), updateData)
		if err != nil {
			log.Println("updateplant : ", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "could not update plant",
			})
			return
		}

		c.JSON(http.StatusOK, response)
	}
}
