package routes

import (
	"conductor/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(router *gin.Engine, db *gorm.DB) {
	conductorController := controllers.ConductorController{DB: db}

	api := router.Group("/api")
	{
		api.POST("/conductores", conductorController.CreateConductor)
		api.GET("/conductores", conductorController.GetConductores)
		api.GET("/conductores/:id", conductorController.GetConductorByID)
		api.PUT("/conductores/:id", conductorController.UpdateConductor)
		api.DELETE("/conductores/:id", conductorController.DeleteConductor)
	}
}
