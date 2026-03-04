package handlers

import (
	"mvc-gorm/repositories"
	"strconv"

	"github.com/gofiber/fiber/v2"
)


type CategoryHanlder struct {
	repo *repositories.CategoryRepository
}

func NewCategoryHandler(repo *repositories.CategoryRepository)*CategoryHanlder{
	return &CategoryHanlder{repo: repo}
}

func (h *CategoryHanlder) GetAll(c *fiber.Ctx) error {
	categories, err := h.repo.GetAll()
	if err !=nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success" 	: false,
			"message" 	: "fetching nya gagal bro coba cek lagi",
			"error"		: err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"succes"	: true,
		"data"		: categories,
	})
}

func (h *CategoryHanlder) GetByID(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10,32)
	if err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success" : false,
			"message" : "Invalid ID",
		})
	}

	category, err := h.repo.GetByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success" : false,
			"message" : "category gada blog !!",
		})
	}

	return c.JSON(fiber.Map{
		"success" : true,
		"data" : category,
	})
}