package models

import "gorm.io/gorm"

type Conductor struct {
	ID                uint   `json:"id" gorm:"primaryKey"`
	Nombre            string `json:"nombre"`
	Descripcion       string `json:"descripcion"`
	CorreoElectronico string `json:"correo_electronico"`
	Telefono          string `json:"telefono"`
	Edad              uint   `json:"edad"`
	RegistroCocheID   uint   `json:"registro_coche_id"` // solo guardamos el ID
}

// Inicializa la tabla en la base de datos
func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&Conductor{})
}
