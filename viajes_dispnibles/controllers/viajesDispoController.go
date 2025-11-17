package controllers

import (
	"marco/api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ViajesDipoController struct {
	DB *gorm.DB
}

// Crear un nuevo viaje disponible
func (cc *ViajesDipoController) CreateViajesDisponibles(c *gin.Context) {
	var vd models.ViajesDisponibles
	if err := c.ShouldBindJSON(&vd); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	cc.DB.Create(&vd)
	c.JSON(http.StatusCreated, vd)
}

// Obtener todos los viaje disponible
func (cc *ViajesDipoController) GetViajeDisponible(c *gin.Context) {
	var vd []models.ViajesDisponibles
	cc.DB.Find(&vd)
	c.JSON(http.StatusOK, vd)
}

// Obtener un viaje disponible por ID
func (cc *ViajesDipoController) GetViajeDisponibleByID(c *gin.Context) {
	id := c.Param("id")
	var vd models.ViajesDisponibles
	if err := cc.DB.First(&vd, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "viaje disponible no encontrado"})
		return
	}
	c.JSON(http.StatusOK, vd)
}

// Actualizar viaje disponible
func (cc *ViajesDipoController) UpdateViajeDisponible(c *gin.Context) {
	id := c.Param("id")
	var vd models.ViajesDisponibles
	if err := cc.DB.First(&vd, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "viaje disponible no encontrado"})
		return
	}

	var input models.ViajesDisponibles
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cc.DB.Model(&vd).Updates(input)
	c.JSON(http.StatusOK, vd)
}

// Eliminar viaje disponible
func (cc *ViajesDipoController) DeleteViajeDisponible(c *gin.Context) {
	id := c.Param("id")
	cc.DB.Delete(&models.ViajesDisponibles{}, id)
	c.JSON(http.StatusOK, gin.H{"message": "viaje disponible eliminado"})
}
