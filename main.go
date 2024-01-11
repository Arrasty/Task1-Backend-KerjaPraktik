package main

import (
	"github.com/Arrasty/tugas/database"
	"github.com/Arrasty/tugas/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	_ "github.com/lib/pq"
)

func main() {

	//panggil func connect dari package database untuk inisalisasi koneksi ke database
	database.Connect()

	//buat instance fiber pake logger dan cors
	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New())

	//setup route dari func SetupRoutes dari package router
	router.SetupRoutes(app)

	// handle unavailable route
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})
	app.Listen(":8080") // => eksekusi"
}
