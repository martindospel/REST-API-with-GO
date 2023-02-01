package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/martindospel/REST-API-with-GO.git/database"
	"github.com/martindospel/REST-API-with-GO.git/routes"
)

func welcomeRoute(c *fiber.Ctx) error {
	return c.SendStatus(200)
}

func allRoutes(app *fiber.App) {
	app.Get("/api", welcomeRoute)

	app.Post("/api/users", routes.CreateUser)
	app.Get("/api/users", routes.GetUsers)
	app.Get("/api/users/:id", routes.GetUser)
	app.Put("/api/users/:id", routes.UpdateUser)
	app.Delete("/api/users/:id", routes.DeleteUser)

	app.Post("/api/products", routes.CreateProduct)
	app.Get("/api/products", routes.GetProducts)
	app.Get("/api/products/:id", routes.GetProduct)
	app.Put("/api/products/:id", routes.UpdateProduct)
	// app.Delete("/api/products/:id", routes.DeleteProduct)

	app.Post("api/orders", routes.CreateOrder)
	app.Get("api/orders", routes.GetOrders)
	app.Get("api/orders/:id", routes.GetOrder)
	app.Delete("api/orders/:id", routes.DeleteOrder)
}

func main() {
	database.ConnectDb()
	app := fiber.New()

	allRoutes(app)
	log.Fatal(app.Listen(":3000"))
}
