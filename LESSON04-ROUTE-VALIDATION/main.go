package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	v1handler "mamba.com/route-group/internal/api/v1/handler"
	v2handler "mamba.com/route-group/internal/api/v2/handler"
	"mamba.com/route-group/middleware"
	"mamba.com/route-group/utils"
)

func main() {

	if err := utils.RegisterValidators(); err != nil {
		panic(err)
	}

	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	r := gin.Default()

	go middleware.CleanupClients()

	r.Use(middleware.LoggerMiddleware(), middleware.ApiKeyMiddleware(), middleware.RateLimitingMiddleware())

	v1 := r.Group("/api/v1")
	{
		user := v1.Group("/users")
		{
			userHandlerV1 := v1handler.NewUserHandler()
			user.GET("", userHandlerV1.GetUsersV1)
			user.GET("/:id", userHandlerV1.GetUsersByIdV1)
			user.GET("/admin/:uuid", userHandlerV1.GetUsersByUuidV1)
			user.POST("", userHandlerV1.PostUsersV1)
			user.PUT("/:id", userHandlerV1.PutUsersByIdV1)
			user.DELETE("/:id", userHandlerV1.DeleteUsersByIdV1)
		}

		product := v1.Group("/products")
		{
			productHandlerV1 := v1handler.NewProductHandler()
			product.GET("", productHandlerV1.GetProductsV1)
			product.GET("/:slug", productHandlerV1.GetProductsBySlugV1)
			product.POST("", productHandlerV1.PostProductsV1)
			product.PUT("/:id", productHandlerV1.PutProductsByIdV1)
			product.DELETE("/:id", productHandlerV1.DeleteProductsByIdV1)
		}

		category := v1.Group("/categories")
		{
			categoryHandlerV1 := v1handler.NewCategoryHandler()
			category.GET("/:category", categoryHandlerV1.GetCategoryByCategoryV1)
			category.POST("", categoryHandlerV1.PostCategoriesV1)
		}

		news := v1.Group("/news")
		{
			newsHandlerV1 := v1handler.NewNewsHandler()
			news.GET("", newsHandlerV1.GetNewsV1)
			news.GET("/:slug", middleware.SimpleMiddleware(), newsHandlerV1.GetNewsV1)
			news.POST("", newsHandlerV1.PostNewsV1)
			news.POST("/upload-file", newsHandlerV1.PostUploadFileNewsV1)
			news.POST("/upload-multiple-file", newsHandlerV1.PostUploadMultipleFileNewsV1)
		}

	}

	v2 := r.Group("/api/v2")
	{
		userV2 := v2.Group("/users")
		{
			userHandlerV2 := v2handler.NewUserHandler()
			userV2.GET("", userHandlerV2.GetUsersV2)
			userV2.GET("/:id", userHandlerV2.GetUsersByIdV2)
			userV2.POST("", userHandlerV2.PostUsersV2)
			userV2.PUT("/:id", userHandlerV2.PutUsersByIdV2)
			userV2.DELETE("/:id", userHandlerV2.DeleteUsersByIdV2)
		}

	}

	// False : Ẩn hiện file khi call API, Ex: localhost:8080/images
	r.StaticFS("/images", gin.Dir("./uploads", false))

	r.Run(":8080")
}

// Zerolog lưu lại log dưới dạng JSON kèm theo timestamp
// Lưu lại các thông tin như method, path, querry,... trong log
// Sắp xếp các middleware cho phù hợp Logger -> API -> Ratelimit
