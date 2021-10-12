package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/victorsaavedra/todo-list/models"
)

func SetUpRoutes(app *fiber.App) {
	app.Get("/todos", models.GetTodos)
	app.Get("/todos/:id", models.GetTodoById)
	app.Post("/todos", models.CreateTodos)
	app.Put("/todos/:id", models.UpdateTodo)
	app.Delete("/todos/:id", models.DeleteTodo)
}
