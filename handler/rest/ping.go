package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ApiPing(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
