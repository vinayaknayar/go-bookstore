package configs

import (
	"gorm.io/gorm"
	_ "gorm.io/driver/mysql"
)

var (
	db * gorm.DB
)

func Connect() {
}