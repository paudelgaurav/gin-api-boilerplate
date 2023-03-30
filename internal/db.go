package internal

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Database model
type Database struct {
	*gorm.DB
}

// SetupModels : initializing mysql database
func NewDatabase() Database {
	USER := Config("DB_USER")
	PASS := Config("DB_PASS")
	HOST := Config("DB_HOST")
	PORT := Config("DB_PORT")
	DBNAME := Config("DB_NAME")

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      false,       // Disable color
		},
	)

	createDBDsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", HOST, USER, PASS, DBNAME, PORT)
	database, err := gorm.Open(postgres.Open(createDBDsn), &gorm.Config{Logger: newLogger})

	if err != nil {
		panic(err.Error())
	}

	sqlDB, err := database.DB()
	if err != nil {
		panic(err.Error())
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	return Database{DB: database}

}
