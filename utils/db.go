package utils

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var(
	DB *gorm.DB
	err error
)

func init() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/webapp?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		//Logger:logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err.Error())
	}
}