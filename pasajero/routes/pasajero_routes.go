package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/marco/pasajero/controllers"
	"gorm.io/gorm"
)

func SetupRoutes(router *gin.Engine, db *gorm.DB) {
	pasajeroController := controllers.PasajeroController{DB: db}

	api := router.Group("/api")
	{
		api.POST("/pasajero", pasajeroController.CreatePasajero)
		api.GET("/pasajero", pasajeroController.GetPasajero)
		api.GET("/pasajero/:id", pasajeroController.GetPasajeroByID)
		api.PUT("/pasajero/:id", pasajeroController.UpdatePasajero)
		api.DELETE("/pasajero/:id", pasajeroController.DeletePasajero)
	}
}
