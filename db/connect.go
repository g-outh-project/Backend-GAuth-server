package db

import (
	"fmt"

	"github.com/Backend-GAuth-server/model"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Start() {
	var dbConfig map[string]string
	dbConfig, _ = godotenv.Read()

	mysqlCredentials := fmt.Sprintf(
		"%s:%s@%s(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbConfig["MYSQL_USER"],
		dbConfig["MYSQL_PASSWORD"],
		dbConfig["MYSQL_PROTOCOL"],
		dbConfig["MYSQL_HOST"],
		dbConfig["MYSQL_PORT"],
		dbConfig["MYSQL_DBNAME"],
	)

	db, _ = gorm.Open(mysql.Open(mysqlCredentials), &gorm.Config{})

	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Client{})
}

func GetDB() *gorm.DB {
	return db
}

func CloseDB() {
	database, _ := db.DB()
	database.Close()
}
