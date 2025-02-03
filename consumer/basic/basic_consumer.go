package main

import (
	"context"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{"localhost:9092"},
		Topic:     "example-topic",
		Partition: 0,
		MinBytes:  10e3, // 10KB
		MaxBytes:  10e6, // 10MB
	})

	defer reader.Close()

	for {
		message, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatal("failed to read messages:", err)
		}
		log.Printf("Consumed: %s\n", string(message.Value))
		time.Sleep(5 * time.Second)
	}
}
