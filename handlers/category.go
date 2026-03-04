package handlers

import (
	"mvc-gorm/models"
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

func (h *CategoryHanlder) Create(c *fiber.Ctx) error {
	category := new(models.Category)
	if err := c.BodyParser(category); err !=nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success" 	: false,
			"message" 	: "invalid request body",
			"error"		: err.Error(),
		})	
	}

	if category.Name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success" : false,
			"message" : "field nama gaboleh kosong woy !",
		})
	}

	if err := h.repo.Create(category); err !=nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success" 	: false,
			"message" 	: "gagal untuk membuat kategori",
			"error"		: err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success" 	: true,
		"message" 	: "succes, data berhasil dibuat !",
		"data"		: category,
	})
}

func (h *CategoryHanlder) Update(c *fiber.Ctx) error{
	id, err := strconv.ParseUint(c.Params("id"),10,32)
	if err!=nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success" : false,
			"message" : "invalid ID mas",
		})
	}

	var body map[string]interface{}
	if err := c.BodyParser(&body); err !=nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success"	: false,
			"message"	: "invalid request bro",
			"error"		: err.Error(),
		})
	}

	allowedFields := map[string]bool{"name":true}
	updateData := make(map[string]interface{})

	for key, val := range body {
		if allowedFields[key] {
			updateData[key] = val
		}
	}

	if len(updateData) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success" : false,
			"message" : "keknya fieldnya ga valid deh, coba cek lagi !!",
		})
	}

	category, err := h.repo.Update(uint(id), updateData)
	if err !=nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success" 	: false,
			"message" 	: "gagal menupdate kategori bro",
			"error"		: err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"success" 	: true,
		"message" 	: "kerja bagus, kategori berhasil ter update su !!",
		"data"		: category,
	})
}

func (h *CategoryHanlder) Delete(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err !=nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success" 	: false,
			"message"	: "invalid ID",
		})
	}
	
	category, err := h.repo.GetByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Gagal cek data kategori",
			"error":   err.Error(),
		})
	}

	if category == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Kategori tidak ditemukan",
		})
	}

	if err := h.repo.Delete(uint(id)); err !=nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success" 	: false,
			"message"	: "gagal, coba cek lagi bro !",
			"error"		: err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"success" 	: true,
		"message"	: "sukses, kategori berhasil di delete",
	})
}