package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	topic := "example-topic-1"
	partition := 0

	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	defer conn.Close()

	for {
		message := fmt.Sprintf("Event from P2 at %s", time.Now().Format(time.RFC3339))
		_, err := conn.WriteMessages(
			kafka.Message{Value: []byte(message)},
		)
		if err != nil {
			log.Fatal("failed to write messages:", err)
		}
		fmt.Println("Produced:", message)
		time.Sleep(60 * time.Second) // generate an event every 60 seconds
	}
}
