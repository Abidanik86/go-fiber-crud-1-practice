package main

import (
	"go-fiber-crud/database"
	"go-fiber-crud/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDatabase()

	app := fiber.New()

	// Routes
	app.Get("/books", routes.GetBooks)
	app.Get("/books/:id", routes.GetBook)
	app.Post("/books", routes.CreateBook)
	app.Put("/books/:id", routes.UpdateBook)
	app.Delete("/books/:id", routes.DeleteBook)

	app.Listen(":8080") // সার্ভার চালু করুন
}
