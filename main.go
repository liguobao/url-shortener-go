package main

import (
	"fmt"
	"log"
	service "url-shortener-go/service"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/:shortCode", service.HandleLinkCode)
	r.POST("share-link/", service.HandleCreateLink)
	r.POST("/healthy/readiness", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data": "success",
		})
	})
	log.Println(fmt.Sprintf("router register success!."))

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
