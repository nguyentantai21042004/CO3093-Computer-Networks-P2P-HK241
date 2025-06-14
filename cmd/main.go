package main

import (
	"fmt"
	"log"

	"CO3093-Computer-Networks-P2P-HK241/src/config"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load cấu hình từ file config.yml
	cfg, err := config.LoadConfig("config.yml")
	if err != nil {
		log.Fatalf("Error loading config file: %v", err)
	}

	// Khởi tạo Gin router
	r := gin.Default()

	// Định nghĩa route cơ bản
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, Gin with config.yml!",
			"db_host": cfg.Database.Host,
		})
	})

	// Chạy server theo cổng trong file config.yml
	port := fmt.Sprintf(":%d", cfg.Server.Port)
	log.Printf("Server running at http://localhost%s", port)
	r.Run(port)
}
