package routes

import (
	"mvc-gorm/config"
	"mvc-gorm/handlers"
	"mvc-gorm/repositories"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App){
	
	categoryRepo := repositories.NewCategoryRepository(config.DB)

	categoryHandler := handlers.NewCategoryHandler(categoryRepo)

	api := app.Group("/api")

	category := api.Group("/categories")
	category.Get("/", categoryHandler.GetAll)
	category.Get("/:id", categoryHandler.GetByID)
	category.Post("/", categoryHandler.Create)
	category.Patch("/:id", categoryHandler.Update)
	category.Delete("/:id", categoryHandler.Delete)
}