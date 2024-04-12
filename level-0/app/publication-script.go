package main

import (
	"io"
	"log"
	"os"
	"time"

	"github.com/nats-io/stan.go"
)

func main() {
	sc, err := stan.Connect("demo-cluster", "demo-publish", stan.NatsURL("nats://localhost:4222"))
	if err != nil {
		log.Fatal(err)
	}
	defer sc.Close()

	jsonFiles := []string{"model.json", "some_items.json", "err.json",
							"order_01.json", "order_02.json", "order_03.json", "order_04.json"}

	for _, v := range jsonFiles{
		path, _ := os.Getwd()
		path +=  "/data/" + v
		msg, _ := readFile(path)
		log.Println(string(msg))

		err := sc.Publish("OrderChannel", msg)
		if err != nil {
			log.Println("Error publishing message:", err)
		} else {
			log.Printf("Data form %v published successfully\n", v)
		}

		time.Sleep(4 * time.Second)
	}
}


func readFile(Path string) ([]byte, error) {
	file, err := os.Open(Path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	data, err := io.ReadAll(io.Reader(file))
	if err != nil {
		return nil, err
	}
	return data, nil
}

