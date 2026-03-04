package main

import (
	"log"
	"mvc-gorm/config"
	"mvc-gorm/routes"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main(){

	config.ConnectDB();

	app := fiber.New(fiber.Config{
		AppName: "mvc-gorm",
	})

	app.Use(logger.New())

	routes.RegisterRoutes(app);

	app.Get("/test", func(c *fiber.Ctx) error {
    return c.SendString("OK")
	})

	log.Println("server udah nyala  di port 8181")
	app.Listen(":8181")
}