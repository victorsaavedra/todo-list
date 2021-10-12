package migrations

import (
	"github.com/victorsaavedra/todo-list/database/database"
	"github.com/victorsaavedra/todo-list/models"
)

func Migrate() {
	database.DBConn.AutoMigrate(models.Todo{})
}
