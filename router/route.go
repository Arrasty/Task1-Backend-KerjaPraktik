package router

import (
	"github.com/Arrasty/tugas/handler"
	"github.com/gofiber/fiber/v2"
)

// func SetupRoutes = mengonfigurasi dan menetapkan rute-rute HTTP pada aplikasi Fiber
func SetupRoutes(app *fiber.App) {

	// grouping rute-rute terkait pengguna di bawah '/api/user'
	api := app.Group("/api")
	v1 := api.Group("/user")

	// rute http
	v1.Get("/", handler.GetAllUsers)
	v1.Get("/:id", handler.GetSingleUser)
	v1.Get("/nim/:nim", handler.GetSingleUserByNim)
	v1.Post("/", handler.CreateUser)
	v1.Put("/:id", handler.UpdateUser)
	v1.Delete("/:id", handler.DeleteUserByID)
}
