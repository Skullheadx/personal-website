package main

import (
	"context"
	"embed"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
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

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router.Handler(),
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Server Shutdown ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Println("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}

func getHealth(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Healthy %v", time.Now())})
}

func getMainPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "Skullheadx",
	})
}
