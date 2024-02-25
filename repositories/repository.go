package repositories

import (
	"fmt"
	"os"

	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	gormv2 "gorm.io/gorm"
)

// helper function for connection string
func NewConnectionString() string {
	var (
		host     = os.Getenv("HOST")
		port     = os.Getenv("DB_PORT")
		user     = os.Getenv("USER")
		password = os.Getenv("PASSWORD")
		dbname   = os.Getenv("DATABASE")
	)
	return fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable", user, password, dbname, host, port)
}

// connect to the underlying database
func Connect() (*gormv2.DB, error) {
	host := os.Getenv("HOST")
	dbString := NewConnectionString()

	dsn := fmt.Sprintf("host=%s %s", host, dbString)
	db, err := gormv2.Open(postgres.Open(dsn))
	if err != nil {
		log.Error().Msgf("unable to open database connection firing panic %s", err)
		panic(err)
	}
	return db, nil
}
