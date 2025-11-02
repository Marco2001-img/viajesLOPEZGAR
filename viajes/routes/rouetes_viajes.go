package routes

import (
	"github.com/marco/apiViajes/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(router *gin.Engine, db *gorm.DB) {
	viajesController := controllers.ViajesController{DB: db}

	api := router.Group("/api")
	{
		api.POST("/viajes", viajesController.CreateViajes)
		api.GET("/viajes", viajesController.GetViajes)
		api.GET("/viajes/:id", viajesController.GetViajesByID)
		api.PUT("/viajes/:id", viajesController.UpdateViajes)
		api.DELETE("/viajes/:id", viajesController.DeleteViaje)
	}
}
