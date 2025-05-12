package http

import "github.com/gin-gonic/gin"

func StartServer() {
	r := gin.Default()

	r.POST("/send", handleSend)

	r.Run(".8080")

}
