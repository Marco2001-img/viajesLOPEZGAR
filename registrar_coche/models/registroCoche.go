package models

import "gorm.io/gorm"

type RegistroCoche struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	MarcaCoche  string `json:"MarcaCoche"`
	Descripcion string `json:"Descripcion"`
}

// Inicializa la tabla en la base de datos
func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&RegistroCoche{})
}
