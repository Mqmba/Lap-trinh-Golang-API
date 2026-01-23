package middleware

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

type Client struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

var (
	mu      sync.Mutex
	clients = make(map[string]*Client)
)

func GetClientIP(ctx *gin.Context) string {
	ip := ctx.ClientIP()

	// Tránh qua lớp Proxy mã hóa thành 1 IP khác
	if ip == "" {
		ip = ctx.Request.RemoteAddr
	}
	return ip
}

func GetRateLimiter(ip string) *rate.Limiter {
	mu.Lock()
	defer mu.Unlock()
	client, exists := clients[ip]
	if !exists {
		// Tối đa 10 tokens, 5 requests sẽ cấp lại trong 1 sec
		limiter := rate.NewLimiter(5, 10) // 5 request/sec, burst 10
		newClient := &Client{limiter, time.Now()}
		clients[ip] = newClient

		log.Printf("A clients[%s] - {limiter: %+v, lastSeen: %s}", ip, newClient.limiter, newClient.lastSeen)

		return limiter
	}

	log.Printf("A clients[%s] - {limiter: %+v, lastSeen: %s}", ip, client.limiter, client.lastSeen)
	client.lastSeen = time.Now()
	return client.limiter
}

func CleanupClients() {
	for {
		time.Sleep(time.Minute)
		mu.Lock()
		for ip, client := range clients {
			// Sau 3 phút, gỡ client ra
			if time.Since(client.lastSeen) > 3*time.Minute {
				delete(clients, ip)
			}
		}
		mu.Unlock()
	}
}

// Chạy Apache BenchMark (ab) trong cmd
// ab -n 20 -c 1 -H "X-API-Key:f8c3942e-ecfc-4aca-9a6f-f6c2242e2e72" localhost:8080/api/v1/categories/php
func RateLimitingMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ip := GetClientIP(ctx)
		log.Println("IP Address: ", ip) // ::1 <=> 127.0.0.1

		limiter := GetRateLimiter(ip)

		if !limiter.Allow() {
			ctx.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error":   "Too many requests",
				"message": "Bạn đã gửi quá nhiều request. Hãy thử lại sau",
			})
			return
		}
		ctx.Next()
	}
}
