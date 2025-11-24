package main

import (
	"embed"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

//go:embed assets/* templates/*
var f embed.FS

func main() {
	router := gin.Default()

	templ := template.Must(template.New("").ParseFS(f, "templates/*.tmpl"))
	router.SetHTMLTemplate(templ)

	router.StaticFS("/public", http.FS(f))

	router.StaticFileFS("/favicon.ico", "assets/favicon.ico", http.FS(f))
	router.GET("/health", getHealth)
	router.GET("/", getMainPage)

	if err := router.Run("localhost:8080"); err != nil {
		log.Fatal(err)
	}
}

func getHealth(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Healthy %v", time.Now())})
}

func getMainPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "Main website",
	})
}
