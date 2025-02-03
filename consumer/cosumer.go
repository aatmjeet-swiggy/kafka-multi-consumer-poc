package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	// "github.com/swiggy-private/ads-sos/ads-common/pkg/kafka"
	"github.com/swiggy-private/kafka-client-go/swgykafka"
)

func main() {
	primary, err := swgykafka.NewClusterBuilder("localhost:9092").AuthMechanism(swgykafka.AuthNone).Build()
	if err != nil {
		fmt.Printf("error occurred in creating cluster: %v", err)
		return
	}

	// Build Topics
	topic1, err := swgykafka.NewTopicBuilder("example-topic").Build()
	if err != nil {
		fmt.Printf("error occurred in creating topic1: %v", err)
		return
	}
	topic2, err := swgykafka.NewTopicBuilder("example-topic-1").Build()
	if err != nil {
		fmt.Printf("error occurred in creating topic2: %v", err)
		return
	}

	// Create a map of topics
	topics := map[string]*swgykafka.Topic{
		"example-topic":   topic1,
		"example-topic-1": topic2,
	}

	// Build Consumer config with multiple topics
	consumerConfig, err := swgykafka.NewConsumerConfigBuilder(primary, nil, "client-id", topic1, "cg-id").
		Topics(topics).
		Build()
	if err != nil {
		fmt.Printf("error occurred in creating consumer config: %v", err)
		return
	}

	// Create Consumer Instance
	consumer, err := swgykafka.NewConsumer(*consumerConfig, MyHandler{}, nil)
	if err != nil {
		fmt.Printf("error occurred in creating consumer: %v", err)
		return
	}

	// Start the consumer
	err = consumer.Start()
	if err != nil {
		fmt.Printf("error occurred in starting consumer: %v", err)
		return
	}

	// Below code is just for blocking
	shutdownChannel := make(chan os.Signal, 1)
	signal.Notify(shutdownChannel, os.Interrupt, syscall.SIGTERM)
	<-shutdownChannel
}

// MyHandler is a custom handler implementing the ConsumerHandler interface
type MyHandler struct{}

func (m MyHandler) Handle(record *swgykafka.Record) swgykafka.Status {
	fmt.Printf("Consumed message from topic: %s, value: %s\n", record.Topic, string(record.Value))
	return swgykafka.Success
}
