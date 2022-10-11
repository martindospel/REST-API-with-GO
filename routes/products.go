package routes

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/martindospel/REST-API-with-GO.git/database"
	"github.com/martindospel/REST-API-with-GO.git/models"
)

func CreateProduct(c *fiber.Ctx) error {
	var product models.Product
	if err := c.BodyParser(&product); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	database.Database.Db.Create(&product)
	return c.Status(200).JSON(product)
}

func GetProducts(c *fiber.Ctx) error {
	products := []models.Product{}
	database.Database.Db.Find(&products)
	return c.Status(200).JSON(products)
}

func findProduct(id int, product *models.Product) error {
	database.Database.Db.Find(&product, "id = ?", id)
	if product.ID == 0 {
		return errors.New("product does not exist")
	}
	return nil
}

func GetProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var product models.Product
	if err != nil {
		return c.Status(400).JSON("Id must be an integer")
	}
	if err := findProduct(id, &product); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	return c.Status(200).JSON(product)
}

func UpdateProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var product models.Product
	if err != nil {
		return c.Status(400).JSON("Id must be an integer")
	}
	if err := findProduct(id, &product); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	type UpdateProduct struct {
		Name         string `json:"name"`
		SerialNumber string `json:"serial_number"`
	}
	var updateData UpdateProduct
	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(500).JSON(err.Error())
	}
	product.Name = updateData.Name
	product.SerialNumber = updateData.SerialNumber
	database.Database.Db.Save(&product)
	return c.Status(200).JSON(product)
}

func DeleteProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var product models.Product
	if err != nil {
		return c.Status(400).JSON("Id must be an integer")
	}
	if err := findProduct(id, &product); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	if err := database.Database.Db.Delete(&product).Error; err != nil {
		return c.Status(404).JSON(err.Error())
	}
	return c.Status(200).SendString("Successfully deleted product")
}
