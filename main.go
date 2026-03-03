package main

import (
	"log"
	"mvc-gorm/config"
	"mvc-gorm/routes"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main(){
	config.ConnectDB();
	routes.RegisterRoutes();

	log.Println("server udah nyala  di port 8181")
	http.ListenAndServe(":8181", nil)
}