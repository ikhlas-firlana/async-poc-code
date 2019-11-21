package main

import (
	"log"
	"time"
)

func main() {
	go func() {
		time.Sleep(5 * time.Second)
		log.Printf("TIMER COMES OUT!")
	}()
}
