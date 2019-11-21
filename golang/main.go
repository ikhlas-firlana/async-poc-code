package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	// set send default
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {

		go func() {
			time.Sleep(5 * time.Second)
			log.Printf("TIMER COMES OUT!")
		}()
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
