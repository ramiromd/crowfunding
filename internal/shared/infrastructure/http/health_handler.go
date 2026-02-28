package http

import (
	"time"
	nethttp "net/http"

	"github.com/gin-gonic/gin"
)

func HealthHandler(c *gin.Context) {
	c.JSON(nethttp.StatusOK, gin.H{
		"status": "ok",
		"timestamp": time.Now().UTC(),
	})
}