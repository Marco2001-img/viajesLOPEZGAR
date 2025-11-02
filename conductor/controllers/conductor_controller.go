package controllers

import (
	"conductor/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ConductorController struct {
	DB *gorm.DB
}

// Crear un nuevo conductor
func (cc *ConductorController) CreateConductor(c *gin.Context) {
	var conductor models.Conductor
	if err := c.ShouldBindJSON(&conductor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	cc.DB.Create(&conductor)
	c.JSON(http.StatusCreated, conductor)
}

// Obtener todos los conductores
func (cc *ConductorController) GetConductores(c *gin.Context) {
	var conductores []models.Conductor
	cc.DB.Find(&conductores)
	c.JSON(http.StatusOK, conductores)
}

// Obtener un conductor por ID
func (cc *ConductorController) GetConductorByID(c *gin.Context) {
	id := c.Param("id")
	var conductor models.Conductor
	if err := cc.DB.First(&conductor, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Conductor no encontrado"})
		return
	}
	c.JSON(http.StatusOK, conductor)
}

// Actualizar conductor
func (cc *ConductorController) UpdateConductor(c *gin.Context) {
	id := c.Param("id")
	var conductor models.Conductor
	if err := cc.DB.First(&conductor, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Conductor no encontrado"})
		return
	}

	var input models.Conductor
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cc.DB.Model(&conductor).Updates(input)
	c.JSON(http.StatusOK, conductor)
}

// Eliminar conductor
func (cc *ConductorController) DeleteConductor(c *gin.Context) {
	id := c.Param("id")
	cc.DB.Delete(&models.Conductor{}, id)
	c.JSON(http.StatusOK, gin.H{"message": "Conductor eliminado"})
}
