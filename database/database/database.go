package database

import (
	"github.com/victorsaavedra/todo-list/framework"
	"github.com/victorsaavedra/todo-list/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DBConn *gorm.DB
)

func InitDatabase() {
	var err error
	dbUser := framework.GoDotEnvVariable("DB_USER")
	dbPassword := framework.GoDotEnvVariable("DB_PASSWORD")
	dbHost := framework.GoDotEnvVariable("DB_HOST")
	dbPort := framework.GoDotEnvVariable("DB_PORT")
	dbDatabase := framework.GoDotEnvVariable("DB_DATABASE")

	dsn := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbDatabase
	DBConn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to DB")
	}

	DBConn.AutoMigrate(&models.Todo{})
}
