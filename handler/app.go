package handler

import (
	"toko-belanja-app/infra/config"
	"toko-belanja-app/infra/database"

	"toko-belanja-app/repository/category_repository/category_pg"
	"toko-belanja-app/repository/product_repository/product_pg"
	"toko-belanja-app/repository/transaction_history_repository/transaction_history_pg"
	"toko-belanja-app/repository/user_repository/user_pg"

	"toko-belanja-app/service/auth_service"
	"toko-belanja-app/service/category_service"
	"toko-belanja-app/service/product_service"
	"toko-belanja-app/service/transaction_history_service"
	"toko-belanja-app/service/user_service"

	_ "toko-belanja-app/docs"

	"github.com/gin-gonic/gin"
	
	swaggoFile "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Toko Belanja
// @version 1.0
// description Final Project 4 Kampus Merdeka

// @contact.name GLNG-KS07 - Group 5
// @contact.url https://github.com/yusrililhm/final-project-4-toko-belanja

// @host localhost:8080
// @BasePath /

func StartApplication() {

	// load .env file
	config.LoadEnv()

	// database init and get instance
	database.InitializeDatabase()
	db := database.GetInstanceDatabaseConnection()

	// dependency injection
	userRepo := user_pg.NewUserPg(db)
	userService := user_service.NewUserService(userRepo)
	userHandler := NewUserHandler(userService)

	categoryRepo := category_pg.NewCategoryPg(db)
	categoryService := category_service.NewCategoryService(categoryRepo)
	categoryHandler := NewCategoryHandler(categoryService)

	productRepo := product_pg.NewProductPg(db)
	productService := product_service.NewProductService(productRepo, categoryRepo)
	productHandler := NewProductHandler(productService)

	transactionHistoryRepo := transaction_history_pg.NewTransactionHistoryPg(db)
	transactionHistoryService := transaction_history_service.NewTransactionHistoryService(transactionHistoryRepo, productRepo, userRepo)
	transactionHistoryHandler := NewTransactionHistoryHandler(transactionHistoryService)

	authService := auth_service.NewAuthService(userRepo)

	app := gin.Default()

	// swagger
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggoFile.Handler))

	// routing
	users := app.Group("users")

	{
		users.POST("/register", userHandler.UserRegister)
		users.POST("/login", userHandler.UserLogin)
		users.PATCH("/topup", authService.Authentication(), userHandler.UserTopUp)
	}

	products := app.Group("products")

	{
		products.POST("", authService.Authentication(), authService.AdminAuthorization(), productHandler.AddProduct)
		products.GET("", authService.Authentication(), productHandler.GetProducts)
		products.PUT("/:productId", authService.Authentication(), authService.AdminAuthorization(), productHandler.UpdateProduct)
		products.DELETE("/:productId", authService.Authentication(), authService.AdminAuthorization(), productHandler.DeleteProduct)
	}

	categories := app.Group("categories")

	{
		categories.POST("", authService.Authentication(), authService.AdminAuthorization(), categoryHandler.AddCategory)
		categories.GET("", authService.Authentication(), authService.AdminAuthorization(), categoryHandler.GetCategories)
		categories.PATCH("/:categoryId", authService.Authentication(), authService.AdminAuthorization(), categoryHandler.UpdateCategory)
		categories.DELETE("/:categoryId", authService.Authentication(), authService.AdminAuthorization(), categoryHandler.DeleteCategory)
	}

	transactionHistories := app.Group("transactions")

	{
		transactionHistories.POST("", authService.Authentication(), transactionHistoryHandler.AddTransaction)
		transactionHistories.GET("/my-transactions", authService.Authentication(), transactionHistoryHandler.GetMyTransaction)
		transactionHistories.GET("/user-transactions", authService.Authentication(), authService.AdminAuthorization(), transactionHistoryHandler.GetUsersTransaction)
	}

	app.Run(":" + config.AppConfig().Port)
}
