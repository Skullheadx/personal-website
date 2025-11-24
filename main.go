package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/health", getHealth)

	if err := router.Run("localhost:8080"); err != nil {
		log.Printf("Router error: %v", err)
	}
}

func getHealth(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Healthy %v", time.Now())})
}
