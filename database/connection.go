package database

import (
	"fmt"
	"log"
	"os"
	"vcm-medical-platform/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect() error {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	sslmode := os.Getenv("DB_SSLMODE")

	if host == "" {
		host = "localhost"
	}
	if port == "" {
		port = "5432"
	}
	if sslmode == "" {
		sslmode = "disable"
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", 
		host, port, user, password, dbname, sslmode)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	log.Println("✅ Database connection established")
	return nil
}

func Migrate() error {
	if DB == nil {
		return fmt.Errorf("database connection not established")
	}

	// Auto migrate the schema
	err := DB.AutoMigrate(
		&models.UserType{},
		&models.User{},
		&models.Country{},
		&models.State{},
		&models.City{},
		&models.District{},
	)

	if err != nil {
		return fmt.Errorf("failed to migrate database: %w", err)
	}

	log.Println("✅ Database migration completed")
	return nil
}

func SeedData() error {
	if DB == nil {
		return fmt.Errorf("database connection not established")
	}

	// Seed user types
	userTypes := []models.UserType{
		{UserType: 0, UserTypeName: "Client/Patient"},
		{UserType: 1, UserTypeName: "Agent"},
		{UserType: 2, UserTypeName: "Sales Channel"},
		{UserType: 3, UserTypeName: "Influencer"},
		{UserType: 4, UserTypeName: "Distributor"},
		{UserType: 5, UserTypeName: "Doctor"},
		{UserType: 10, UserTypeName: "Operator"},
		{UserType: 11, UserTypeName: "Admin"},
		{UserType: 12, UserTypeName: "Super Admin"},
	}

	for _, userType := range userTypes {
		var existing models.UserType
		result := DB.Where("usertype = ?", userType.UserType).First(&existing)
		if result.Error == gorm.ErrRecordNotFound {
			if err := DB.Create(&userType).Error; err != nil {
				log.Printf("Error creating user type %d: %v", userType.UserType, err)
			}
		}
	}

	// Seed basic countries
	countries := []models.Country{
		{CdCountry: 1, CountryName: "United States", CountryAbbr: "US"},
		{CdCountry: 86, CountryName: "China", CountryAbbr: "CN"},
		{CdCountry: 44, CountryName: "United Kingdom", CountryAbbr: "UK"},
		{CdCountry: 33, CountryName: "France", CountryAbbr: "FR"},
		{CdCountry: 49, CountryName: "Germany", CountryAbbr: "DE"},
		{CdCountry: 81, CountryName: "Japan", CountryAbbr: "JP"},
		{CdCountry: 82, CountryName: "South Korea", CountryAbbr: "KR"},
		{CdCountry: 91, CountryName: "India", CountryAbbr: "IN"},
	}

	for _, country := range countries {
		var existing models.Country
		result := DB.Where("cd_country = ?", country.CdCountry).First(&existing)
		if result.Error == gorm.ErrRecordNotFound {
			if err := DB.Create(&country).Error; err != nil {
				log.Printf("Error creating country %d: %v", country.CdCountry, err)
			}
		}
	}

	log.Println("✅ Database seeding completed")
	return nil
}
