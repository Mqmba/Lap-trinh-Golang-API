package main

import (
	"github.com/gin-gonic/gin"
	v1handler "mamba.com/route-group/internal/api/v1/handler"
	v2handler "mamba.com/route-group/internal/api/v2/handler"
)

func main() {
	r := gin.Default()

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
		}

		news := v1.Group("/news")
		{
			newsHandlerV1 := v1handler.NewNewsHandler()
			news.GET("", newsHandlerV1.GetNewsV1)
			news.GET("/:slug", newsHandlerV1.GetNewsV1)
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

	r.Run(":8080")
}
