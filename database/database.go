package database

import (
	"fmt"
	"os"

	"github.com/mofodox/TILNotes/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Global DB pointer
var DB *gorm.DB

func Connect() error {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", os.Getenv("DBHost"), os.Getenv("DBUser"), os.Getenv("DBPassword"), os.Getenv("DBName"), os.Getenv("DBPort"))
	
	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	DB = connection

	connection.Migrator().DropTable(&models.Note{})
	connection.AutoMigrate(&models.Note{})

	return nil
}
