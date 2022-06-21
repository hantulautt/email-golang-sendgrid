package config

import (
	"email/exception"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMysqlDatabase(configuration Config) *gorm.DB {
	dsn := configuration.Get("MYSQL_URI")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// if there is an error opening the connection, handle it
	exception.PanicIfNeeded(err)

	return db
}
