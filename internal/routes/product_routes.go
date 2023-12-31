package routes

import (
	"Tubes-EAI/internal/controllers"
	"Tubes-EAI/internal/middleware"
	"Tubes-EAI/internal/services"
	"github.com/gofiber/fiber/v2"
)

func SetUpProductRoutes(router fiber.Router, productServices services.ProductService) {
	productController := controllers.NewProductController(productServices)

	// Categories
	categories := router.Group("/categories")
	categories.Get("", productController.GetCategories)

	// Products
	products := router.Group("/products").Use(middleware.AdminAuthentication(middleware.AuthConfig{
		Unauthorized: func(ctx *fiber.Ctx) error {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		},
	}))

	products.Post("", productController.AddProduct)
	products.Put("/:id", productController.UpdateProduct)
	products.Delete("/:id", productController.DeleteProduct)

	// Global Search
	serach := router.Group("/search")
	serach.Get("/products", productController.GetAllProducts)
	serach.Get("/product", productController.FindProduct)

}
