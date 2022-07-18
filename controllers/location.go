package controllers

import (
	"go-auth/database"
	"go-auth/models"

	"github.com/gofiber/fiber/v2"
)

func countries(c *fiber.Ctx) error {
	var province models.Countries

	database.DB.Find(&province)

	return c.JSON(province)
}

// func Province(c *fiber.Ctx) error {
// 	var province models.Provinces

// 	Conn.Preload("Country_ID").Find(&province)
// }

// func Cities() {
// 	var city models.Cities

// 	Conn.Preload("Country_ID").Find(&city)
// }

// func Districts() {
// 	var district models.Districts

// 	Conn.Preload("Country_ID").Find(&district)
// }

// func SubDistrict() {
// 	var subdistrict models.SubDistrict

// 	Conn.Preload("Country_ID").Find(&subdistrict)
// }
