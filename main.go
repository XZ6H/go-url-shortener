package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/xz6h/go-url-shortener/handler"
	"github.com/xz6h/go-url-shortener/store"
)

func main() {
	r := gin.Default()
	r.GET("/", func (c *gin.Context)  {
		c.JSON(200, gin.H{
			"message": "Go url short",
		})
	})

	r.POST("/create-short-url", func(ctx *gin.Context) {
		handler.CreateShortUrl(ctx)
	})

	r.GET("/:shortUrl", func(ctx *gin.Context) {
		handler.HandleShortUrlRedirect(ctx);
	})

	store.InitializeStore()
	
	err := r.Run(":9808")
	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server - Error: %v", err))
	}
}
