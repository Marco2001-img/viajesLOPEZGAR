package main

import (
	"fmt"
	"log"
	"os"

	"github.com/marco/pasajero/models"
	"github.com/marco/pasajero/routes"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Variables de entorno se mantienen igual
	host := os.Getenv("POSTGRES_HOST")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")
	port := os.Getenv("POSTGRES_PORT")

	// CORRECCIÓN: Se cambia TimeZone=America/Lima a TimeZone=UTC
	// Esto asegura que la conexión sea estándar y universal, resolviendo el error.
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC", // <-- ¡CORREGIDO AQUÍ!
		host, user, password, dbname, port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		// El log.Fatal se ejecuta en la línea 47 de tu archivo (main.go:47 en el log de error)
		log.Fatal("[error] failed to initialize database, got error ", err)
	}

	if err := models.Migrate(db); err != nil {
		log.Fatal("Error al migrar base de datos:", err)
	}

	router := gin.Default()
	routes.SetupRoutes(router, db)

	router.Run(":8081")
}
