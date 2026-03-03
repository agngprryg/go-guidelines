package routes

import (
	"mvc-gorm/controllers"
	"net/http"
)

func RegisterRoutes(){
	http.HandleFunc("/products", controllers.GetProduct)
	http.HandleFunc("/products/create", controllers.CreateProduct)
	http.HandleFunc("/products/detail", controllers.GetProductByID)
	http.HandleFunc("/products/update", controllers.UpdateProduct)
	http.HandleFunc("/products/delete", controllers.DeleteProduct)

	http.HandleFunc("/categories", controllers.GetCategory)
	http.HandleFunc("/categories/create", controllers.CreateCategory)
	http.HandleFunc("/categories/detail", controllers.GetCategoryByID)
}