package main

import (
	"log"
	
	crowhttp "github.com/ramiromd/crowfunding/internal/shared/infrastructure/http"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	server.GET("/health", crowhttp.HealthHandler)
	status := server.Run(":8080")
	if (status != nil) {
		log.Fatal("Failed to start server: ", status)
	}

	log.Println("Server started on port 8080")
}