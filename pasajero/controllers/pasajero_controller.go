package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marco/pasajero/models"
	"gorm.io/gorm"
)

type PasajeroController struct {
	DB *gorm.DB
}

func (cc *PasajeroController) CreatePasajero(c *gin.Context) {
	var pasajero models.Pasajero
	if err := c.ShouldBindJSON(&pasajero); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	cc.DB.Create(&pasajero)
	c.JSON(http.StatusCreated, pasajero)
}

func (cc *PasajeroController) GetPasajero(c *gin.Context) {
	var pasajero []models.Pasajero
	cc.DB.Find(&pasajero)
	c.JSON(http.StatusOK, pasajero)
}

func (cc *PasajeroController) GetPasajeroByID(c *gin.Context) {
	id := c.Param("id")
	var pasajero models.Pasajero
	if err := cc.DB.First(&pasajero, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "pasajero no encontrado"})
		return
	}
	c.JSON(http.StatusOK, pasajero)
}

func (cc *PasajeroController) UpdatePasajero(c *gin.Context) {
	id := c.Param("id")
	var pasajero models.Pasajero
	if err := cc.DB.First(&pasajero, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "pasajero no encontrado"})
		return
	}

	var input models.Pasajero
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cc.DB.Model(&pasajero).Updates(input)
	c.JSON(http.StatusOK, pasajero)
}

// Eliminar conductor
func (cc *PasajeroController) DeletePasajero(c *gin.Context) {
	id := c.Param("id")
	cc.DB.Delete(&models.Pasajero{}, id)
	c.JSON(http.StatusOK, gin.H{"message": "pasajero eliminado"})
}
