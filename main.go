package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
)

func main() {
	isProduction := false
	if os.Getenv("PRODUCTION") == "TRUE" {
		fmt.Print("Using Production Configuration\n")
		isProduction = true
		// gin.SetMode(gin.ReleaseMode)
	}
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	router.GET("/health", getHealth)
	router.GET("/", getMainPage)
	// router.GET("/about", getAboutPage)
	router.StaticFile("/favicon.ico", "./assets/favicon.ico")
	router.StaticFile("/assets/resume.pdf", "./assets/Andrew_Montgomery_Resume.pdf")
	if isProduction {
		log.Fatal(autotls.Run(router, "andrew-montgomery.dev"))
	} else {
		log.Fatal(router.Run(":8080"))
	}
}

func getHealth(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Healthy %v", time.Now())})
}

func getMainPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}
