package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/victorsaavedra/todo-list/database"
	"github.com/victorsaavedra/todo-list/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func initDatabase() {
	var err error
	dsn := "mysql_go:password@tcp(db:3306)/todoGo"
	database.DBConn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to DB")
	}
	fmt.Println("DB connected")
	database.DBConn.AutoMigrate(&models.Todo{})
	fmt.Println("Migrated DB")
}

func setUpRoutes(app *fiber.App) {
	app.Get("/todos", models.GetTodos)
}

func helloWorld(c *fiber.Ctx) error {
	return c.SendString("Hello world")
}

func main() {
	app := fiber.New()
	initDatabase()
	app.Get("/", helloWorld)
	setUpRoutes(app)
	app.Listen(":8000")
}