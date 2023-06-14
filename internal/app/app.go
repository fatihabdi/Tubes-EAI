package app

import (
	"os"

	"Tubes-EAI/internal/config"
	"Tubes-EAI/internal/repositories"
	"Tubes-EAI/internal/routes"
	"Tubes-EAI/internal/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func StartApplication() {

	// initialize gin
	app := fiber.New()

	// initialize db
	db, err := config.Connect()
	if err != nil {
		panic(err)
	}

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET, POST, PUT, DELETE",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowCredentials: true,
	}))

	// initialize user repositories
	userRepository := repositories.NewUserRepository(db)

	// initialize user services
	userService := services.NewUserService(userRepository)

	// initialize user routes
	apiEndpoint := app.Group("/api")
	routes.SetUpUserRoutes(apiEndpoint, userService)

	// initialize product repositories
	productRepository := repositories.NewProductsRepository(db)

	// initialize product services
	productService := services.NewProductService(productRepository)

	// initialize product routes
	routes.SetUpProductRoutes(apiEndpoint, productService)

	// start the server
	err = app.Listen(":" + os.Getenv("PORT"))

	if err != nil {
		panic(err)
	}
}
