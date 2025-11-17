package models

import "gorm.io/gorm"

type ViajesDisponibles struct {
	ID          uint `json:"id" gorm:"primaryKey"`
	ConductorID uint `json:"conductor_id"`
	ViajesID    uint `json:"Viajes_id"`
	PasajeroID  uint `json:"Pasajero_id"`
}

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&ViajesDisponibles{})
}
