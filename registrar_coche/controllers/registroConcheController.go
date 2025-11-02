package controllers

import (
	"net/http"
	"registrar_conche/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RegistroCocheController struct {
	DB *gorm.DB
}

// Crear un nuevo conductor
func (cc *RegistroCocheController) CreateCocheRegistro(c *gin.Context) {
	var rc models.RegistroCoche
	if err := c.ShouldBindJSON(&rc); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	cc.DB.Create(&rc)
	c.JSON(http.StatusCreated, rc)
}

// Obtener todos los conductores
func (cc *RegistroCocheController) GetCocheRegistro(c *gin.Context) {
	var rc []models.RegistroCoche
	cc.DB.Find(&rc)
	c.JSON(http.StatusOK, rc)
}

// Obtener un conductor por ID
func (cc *RegistroCocheController) GetConductorCocheByID(c *gin.Context) {
	id := c.Param("id")
	var rc models.RegistroCoche
	if err := cc.DB.First(&rc, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Conductor de coche no encontrado"})
		return
	}
	c.JSON(http.StatusOK, rc)
}

// Actualizar conductor
func (cc *RegistroCocheController) UpdateConductorCoche(c *gin.Context) {
	id := c.Param("id")
	var rc models.RegistroCoche
	if err := cc.DB.First(&rc, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Conductor de coche no encontrado"})
		return
	}

	var input models.RegistroCoche
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cc.DB.Model(&rc).Updates(input)
	c.JSON(http.StatusOK, rc)
}

// Eliminar conductor
func (cc *RegistroCocheController) DeleteConductorCoche(c *gin.Context) {
	id := c.Param("id")
	cc.DB.Delete(&models.RegistroCoche{}, id)
	c.JSON(http.StatusOK, gin.H{"message": "Conductor de coche eliminado"})
}
