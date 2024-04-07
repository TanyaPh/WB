package main

import (
	"log"
	"time"

	"github.com/nats-io/stan.go"
)

func main() {
	// Подключение к NATS Streaming
	sc, err := stan.Connect("demo-cluster", "demo-publish", stan.NatsURL("nats://localhost:1234"))
	if err != nil {
		log.Fatal(err)
	}
	defer sc.Close()

	for {
		msg := []byte("Your message here")
		err := sc.Publish("new_order", msg)
		if err != nil {
			log.Println("Error publishing message:", err)
		} else {
			log.Println("Message published successfully")
		}

		time.Sleep(4 * time.Second)
	}
}
