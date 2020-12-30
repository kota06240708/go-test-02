package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @description API疎通確認
// @version 1.0
// @Tags ping
// @Summary API疎通確認
// @accept application/x-json-stream
// @Success 200 {object} gin.H {"status": "success"}
// @router /api/v1/ping [get]
func ApiPing(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
