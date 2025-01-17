package main

import (
	"log"

	"github.com/QSOLink/QSOLink/database"
	"github.com/QSOLink/QSOLink/qso"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())
	database.ConnectDB()
	defer database.DB.Close()

	api := app.Group("/api")
	qso.Register(api, database.DB)

	log.Fatal(app.Listen(":5001"))
}
