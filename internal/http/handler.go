package http

import (
	"EmaiSender/internal/emai"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func handleSend(content *gin.Context) {
	var req EmailRequest
	if err := jsonParse(content, &req); err != nil {
		return
	}

	if err := req.Validate(); err != nil {
		content.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	go func(email, name string) {
		err := emai.SendEmail(email, name)
		if err != nil {

			fmt.Println("email error:", err)
		}
	}(req.Email, req.Name)

	content.JSON(http.StatusOK, gin.H{"success": "success with request "})

}
