package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
)

func main() {
	// make amqp
	conn, err := amqp.Dial("amqp://localhost:5672/")
	if err != nil {
		fmt.Errorf("error amq %v", err)
	}
	defer conn.Close()
	ch, err := conn.Channel()
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			time.Sleep(5 * time.Second)
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	go func() {
		// set send default
		r := gin.Default()
		r.GET("/", func(c *gin.Context) {

			q, err := ch.QueueDeclare(
				"hello", // name
				false,   // durable
				false,   // delete when unused
				false,   // exclusive
				false,   // no-wait
				nil,     // arguments
			)
			if err != nil {
				fmt.Errorf("Hello %v", err)
			}

			err = ch.Publish(
				"",     // exchange
				q.Name, // routing key
				false,  // mandatory
				false,  // immediate
				amqp.Publishing{
					ContentType: "text/plain",
					Body:        []byte("TIMER COMES OUT!"),
				})
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})

		r.Run() // listen and serve on 0.0.0.0:8080
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
