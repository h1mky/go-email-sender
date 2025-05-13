package http

import (
	"EmaiSender/internal/emai"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type EmailRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Terms bool   `json:"terms"`
}

func handleSend(content *gin.Context) {
	var req EmailRequest
	if err := content.ShouldBindJSON(&req); err != nil {
		content.JSON(http.StatusBadRequest, gin.H{"error": "not valid json"})
		return
	}

	switch {
	case req.Name == "":
		content.JSON(http.StatusBadRequest, gin.H{"error": "required field"})
		return
	case req.Email == "":
		content.JSON(http.StatusBadRequest, gin.H{"error": "required field"})
		return
	case req.Terms == false:
		content.JSON(http.StatusBadRequest, gin.H{"error": "required field"})
		return
	}

	if !strings.Contains(req.Email, "@") {
		content.JSON(http.StatusBadRequest, gin.H{"error": "wrong email reference"})
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
