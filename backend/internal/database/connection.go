package database

import (
	"vcm-medical-platform/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Connect(databaseURL string) (*gorm.DB, error) {
	config := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}

	db, err := gorm.Open(postgres.Open(databaseURL), config)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.UserType{},
		&models.User{},
		&models.Assessment{},
		&models.Appointment{},
		&models.Order{},
		&models.ChatRoom{},
		&models.ChatMessage{},
	)
}
