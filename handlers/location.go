package handlers

import (
	"vcm-medical-platform/database"
	"vcm-medical-platform/models"

	"github.com/gofiber/fiber/v2"
)

// GetCountries - Get all countries
func GetCountries(c *fiber.Ctx) error {
	var countries []models.Country
	if err := database.DB.Find(&countries).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to fetch countries",
		})
	}

	return c.JSON(fiber.Map{
		"countries": countries,
	})
}

// GetStates - Get states by country
func GetStates(c *fiber.Ctx) error {
	countryID := c.Params("countryId")
	
	var states []models.State
	if err := database.DB.Where("cd_country = ?", countryID).Find(&states).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to fetch states",
		})
	}

	return c.JSON(fiber.Map{
		"states": states,
	})
}

// GetCities - Get cities by country and state
func GetCities(c *fiber.Ctx) error {
	countryID := c.Params("countryId")
	stateID := c.Params("stateId")
	
	var cities []models.City
	if err := database.DB.Where("cd_country = ? AND cd_state = ?", countryID, stateID).Find(&cities).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to fetch cities",
		})
	}

	return c.JSON(fiber.Map{
		"cities": cities,
	})
}

// GetDistricts - Get districts by country, state, and city
func GetDistricts(c *fiber.Ctx) error {
	countryID := c.Params("countryId")
	stateID := c.Params("stateId")
	cityID := c.Params("cityId")
	
	var districts []models.District
	if err := database.DB.Where("cd_country = ? AND cd_state = ? AND cd_city = ?", 
		countryID, stateID, cityID).Find(&districts).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to fetch districts",
		})
	}

	return c.JSON(fiber.Map{
		"districts": districts,
	})
}
