package controllers

import (
	"net/http"

	"github.com/marco/apiViajes/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ViajesController struct {
	DB *gorm.DB
}

// Crear un nuevo conductor
func (cc *ViajesController) CreateViajes(c *gin.Context) {
	var viajes models.Viajes
	if err := c.ShouldBindJSON(&viajes); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	cc.DB.Create(&viajes)
	c.JSON(http.StatusCreated, viajes)
}

// Obtener todos los conductores
func (cc *ViajesController) GetViajes(c *gin.Context) {
	var viajes []models.Viajes
	cc.DB.Find(&viajes)
	c.JSON(http.StatusOK, viajes)
}

// Obtener un conductor por ID
func (cc *ViajesController) GetViajesByID(c *gin.Context) {
	id := c.Param("id")
	var viajes models.Viajes
	if err := cc.DB.First(&viajes, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "viaje no encontrado"})
		return
	}
	c.JSON(http.StatusOK, viajes)
}

// Actualizar conductor
func (cc *ViajesController) UpdateViajes(c *gin.Context) {
	id := c.Param("id")
	var viajes models.Viajes
	if err := cc.DB.First(&viajes, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "viaje no encontrado"})
		return
	}

	var input models.Viajes
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cc.DB.Model(&viajes).Updates(input)
	c.JSON(http.StatusOK, viajes)
}

// Eliminar conductor
func (cc *ViajesController) DeleteViaje(c *gin.Context) {
	id := c.Param("id")
	cc.DB.Delete(&models.Viajes{}, id)
	c.JSON(http.StatusOK, gin.H{"message": "viaje eliminado"})
}
