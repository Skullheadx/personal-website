package main

import (
	"embed"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

//go:embed assets/*
var f embed.FS

func main() {
	router := gin.Default()

	router.StaticFS("/public", http.FS(f))
	router.StaticFileFS("/favicon.ico", "assets/favicon.ico", http.FS(f))

	router.GET("/health", getHealth)

	if err := router.Run("localhost:8080"); err != nil {
		log.Fatal(err)
	}
}

func getHealth(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Healthy %v", time.Now())})
}
