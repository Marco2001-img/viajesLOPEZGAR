package routes

import (
	"registrar_conche/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(router *gin.Engine, db *gorm.DB) {
	RegistroCocheController := controllers.RegistroCocheController{DB: db}

	api := router.Group("/api")
	{
		api.POST("/conductoresCoche", RegistroCocheController.CreateCocheRegistro)
		api.GET("/conductoresCoche", RegistroCocheController.GetCocheRegistro)
		api.GET("/conductoresCoche/:id", RegistroCocheController.GetConductorCocheByID)
		api.PUT("/conductoresCoche/:id", RegistroCocheController.UpdateConductorCoche)
		api.DELETE("/conductoresCoche/:id", RegistroCocheController.DeleteConductorCoche)
	}
}
