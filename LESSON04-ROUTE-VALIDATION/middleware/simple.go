package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
)

func SimpleMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// Trước khi bắt đầu vào Handler (Before)
		log.Println("Start func - Check from Middleware")
		ctx.Writer.Write([]byte("Start func - Check from Middleware"))

		ctx.Next() // Đi vào Handler

		// Sau khi Handler xử lý xong (After)
		log.Println("End func - Check from Middleware")
		ctx.Writer.Write([]byte("End func - Check from Middleware"))
	}
}

// Middleware luôn phải đặt trước Handler trong main.go nếu sử dụng route cụ thể
