package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
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
	if req.Name == "" {
		content.JSON(http.StatusBadRequest, gin.H{"error": "required fields"})
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

	err := mail.SendEmail(req.Email)

	if err != nil {
		content.JSON(http.StatusInternalServerError, gin.H{"error": "error with server"})
		return
	}
	content.JSON(http.StatusOK, gin.H{"success": "success with request "})

}
