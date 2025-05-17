package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func jsonParse(content *gin.Context, req interface{}) error {
	if err := content.ShouldBindJSON(req); err != nil {
		content.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON format"})
		return err
	}
	return nil
}
