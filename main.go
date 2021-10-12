package main

import (
	"github.com/victorsaavedra/todo-list/database/database"
	"github.com/victorsaavedra/todo-list/database/migrations"
	"github.com/victorsaavedra/todo-list/framework"
	"github.com/victorsaavedra/todo-list/routes"
)

func main() {
	app := framework.GetApp()
	database.InitDatabase()
	migrations.Migrate()
	routes.SetUpRoutes(app)

	app.Listen(":8000")
}