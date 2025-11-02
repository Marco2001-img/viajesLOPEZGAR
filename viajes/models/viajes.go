package models

import "gorm.io/gorm"

type Viajes struct {
	ID                 uint   `json:"id" gorm:"primaryKey"`
	CuidadOrigen       string `json:"CuidadOrigen"`
	CuidadDestino      string `json:"CuidadDestino"`
	HoraSalida         string `json:"HoraSalida"`
	HoraLlegada        string `json:"HoraLlegada"`
	DescripcionLlegada string `json:"DescripcionLlegada"`
	DescripcionSalida  string `json:"DescripcionSalida"`
	Precio             uint   `json:"precio"`
}

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&Viajes{})
}
