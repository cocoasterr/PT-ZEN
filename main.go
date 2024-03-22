package main

import (
	"log"

	"github.com/cocoasterr/kalbe_test/controllers"
	PGConfig "github.com/cocoasterr/kalbe_test/infra/db/postgres"
	PGRepository "github.com/cocoasterr/kalbe_test/infra/db/postgres/repository/databasesql"
	"github.com/cocoasterr/kalbe_test/middleware"
	"github.com/cocoasterr/kalbe_test/services"
	"github.com/gofiber/fiber/v2"
)

func main() {
	db_dsn, _ := PGConfig.CreateDatabase("assesment_ptzen")
	db, err := PGConfig.ConnectPG(db_dsn)
	if err != nil {
		log.Fatal("Failed to Connect DB")
	}
	app := fiber.New()

	ProductRepo := PGRepository.NewProductRepository(db)
	ProductService := services.NewProductService(ProductRepo.Repository, ProductRepo)
	ProductController := controllers.NewProductController(*ProductService)

	PesananRepo := PGRepository.NewPesananRepository(db)
	PesananService := services.NewPesananService(PesananRepo.Repository, PesananRepo, ProductService)
	PesananController := controllers.NewPesananController(*PesananService)

	UserRepo := PGRepository.NewUserRepository(db)
	UserService := services.NewUserService(UserRepo.Repository)
	userController := controllers.NewUserController(*UserService)

	// api
	api := app.Group("/api")
	ProductGroup := api.Group("/product")

	ProductGroup.Post("/", middleware.DeserializeUser, ProductController.CreateController)
	ProductGroup.Get("/", middleware.DeserializeUser, ProductController.IndexController)
	ProductGroup.Get("/:id", middleware.DeserializeUser, ProductController.FindByController)
	ProductGroup.Put("/:id", middleware.DeserializeUser, ProductController.UpdateController)
	ProductGroup.Put("/delete/:id", middleware.DeserializeUser, ProductController.DeleteController)

	pesananGroup := api.Group("/pesanan")

	pesananGroup.Post("/", middleware.DeserializeUser, PesananController.CreateController)
	pesananGroup.Get("/", middleware.DeserializeUser, PesananController.IndexController)
	pesananGroup.Get("/:id", middleware.DeserializeUser, PesananController.FindByController)

	authGroup := api.Group("/auth")

	authGroup.Post("/register", userController.RegisterUser)
	authGroup.Post("/login", userController.Login)
	authGroup.Get("/profile", middleware.DeserializeUser, userController.Profile)
	authGroup.Put("/reset-password", middleware.DeserializeUser, userController.ResetPassword)
	authGroup.Post("/logout", middleware.DeserializeUser, userController.Logout)

	log.Fatal(app.Listen(":3000"))
}
