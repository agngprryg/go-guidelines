package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB(){
	dsn := "root:root@tcp(127.0.0.1:8889)/belajar-golang?parseTime=true"

	db,err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	println("Berhasil connect ke database:", db)
}