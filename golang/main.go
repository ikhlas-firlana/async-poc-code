package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	var wg sync.WaitGroup
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		wg.Add(1)
		go func() {
			time.Sleep(5 * time.Second)
			fmt.Println("TIMER COMES OUT!")
			defer wg.Done()
		}()
		c.JSON(200, gin.H{
			"message": "pong",
		})
		wg.Wait() // here part not working.
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
