package models

import (
	"github.com/gofiber/fiber/v2"
	"github.com/victorsaavedra/todo-list/database"
)

type Todo struct {
	ID uint `gorm:"primarykey" json:"id"`
	Title string `json:"title"`
	Completed bool `json:"completed"`
}

func GetTodos(c *fiber.Ctx) error {
	db := database.DBConn
	var todos []Todo
	db.Find(&todos)

	return c.JSON(&todos)
}