package db

import (
	"fmt"

	"github.com/Backend-GAuth-server/utils"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func Start() {
	var dbConfig map[string]string
	dbConfig, err := godotenv.Read()
	utils.HandlePanic(err)

	mysqlCredentials := fmt.Sprintf(
		"%s:%s@%s(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbConfig["MYSQL_USER"],
		dbConfig["MYSQL_PASSWORD"],
		dbConfig["MYSQL_PROTOCOL"],
		dbConfig["MYSQL_HOST"],
		dbConfig["MYSQL_PORT"],
		dbConfig["MYSQL_DBNAME"],
	)

	db, err = gorm.Open(mysql.Open(mysqlCredentials), &gorm.Config{})

	utils.HandlePanic(err)
}

func GetDB() *gorm.DB {
	return db
}

func CloseDB() {
	database, _ := db.DB()
	database.Close()
}
