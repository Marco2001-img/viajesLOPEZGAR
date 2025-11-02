package models

import "gorm.io/gorm"

type Pasajero struct {
	ID                uint   `json:"id" gorm:"primaryKey"`
	Nombre            string `json:"nombre"`
	ApellidoPaterno   string `json:"apellidoPaterno"`
	ApellidoMaterno   string `json:"apellidoMaterno"`
	Edad              uint   `json:"edad"`
	Cuidadorigen      string `json:"Cuidadorigen"`
	TelefonoPasajero  string `json:"telefonoPasajero"`
	CorreoElectronico string `json:"CorreoElectronico"`
}

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&Pasajero{})
}
