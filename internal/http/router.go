package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func StartServer() {
	r := gin.Default()

	r.POST("/send", handleSend)

	fmt.Println("🚀 Сервер запущен и слушает на http://localhost:8080")

	err := r.Run(":8080")
	if err != nil {
		fmt.Println("Ошибка при запуске сервера:", err)

		return
	}

}
