package router

import (
	"fmt"

	"github.com/YakovlevIgA/anotherTest/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {

	api := app.Group("/api", logger.New())

	api.Post("/tasks", handler.AddTask)

	api.Get("/tasks", handler.ShowTasks)

	api.Put("/tasks/:id", handler.EditTask)

	api.Delete("/tasks/:id", handler.DeleteTask)

	fmt.Println("All routes have been registrated!")
}
