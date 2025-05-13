package http

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func StartServer() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // замени на порт фронта
		AllowMethods:     []string{"POST", "GET", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.POST("/send", handleSend)

	fmt.Println("🚀 Сервер запущен и слушает на http://localhost:8080")

	err := r.Run(":8080")
	if err != nil {
		fmt.Println("Ошибка при запуске сервера:", err)

		return
	}

}
