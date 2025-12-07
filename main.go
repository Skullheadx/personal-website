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
	router.LoadHTMLGlob("templates/*.html")

	router.GET("/health", getHealth)
	router.GET("/", getMainPage)
	// router.GET("/about", getAboutPage)
	router.StaticFile("/favicon.ico", "./assets/favicon.ico")
	router.StaticFile("/assets/resume.pdf", "./assets/Andrew_Montgomery_Resume.pdf")
	if err := router.Run("localhost:8080"); err != nil {
		log.Fatal(err)
	}
}

func getHealth(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Healthy %v", time.Now())})
}

func getMainPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}
