package main

import (
	"log"
	
	crowhttp "github.com/ramiromd/crowfunding/internal/shared/infrastructure/http"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	server.GET("/health", crowhttp.HealthHandler)
	if err := server.Run(":8080"); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}