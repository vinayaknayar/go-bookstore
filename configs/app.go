package configs

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

var (
	db * gorm.DB
)

// connect to database
func Connect() (*gorm.DB, error) {
	//Load  environment variables
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// create connection string
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)

	// initailize the mysql session
	d, err := gorm.Open(mysql.Open(connection), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db = d
	return db, nil
}

// get database
func GetDB() *gorm.DB {
	return db
}