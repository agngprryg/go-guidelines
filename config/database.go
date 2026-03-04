package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB(){
	dsn := "root:root@tcp(127.0.0.1:8889)/belajar-golang?parseTime=true&loc=Asia%2FJakarta"

	db,err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	DB = db
	println("Berhasil connect ke database:", db)
}