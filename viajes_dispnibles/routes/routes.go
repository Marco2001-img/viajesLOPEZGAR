package routes

import (
	"marco/api/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(router *gin.Engine, db *gorm.DB) {
	ViajesDispoController := controllers.ViajesDipoController{DB: db}

	api := router.Group("/api")
	{
		api.POST("/viajesDisponibles", ViajesDispoController.CreateViajesDisponibles)
		api.GET("/viajesDisponibles", ViajesDispoController.GetViajeDisponible)
		api.GET("/viajesDisponibles/:id", ViajesDispoController.GetViajeDisponibleByID)
		api.PUT("/viajesDisponibles/:id", ViajesDispoController.UpdateViajeDisponible)
		api.DELETE("/viajesDisponibles/:id", ViajesDispoController.DeleteViajeDisponible)
	}
}
