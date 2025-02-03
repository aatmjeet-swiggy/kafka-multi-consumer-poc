package main

import (
	"fmt"
	"log"
	"time"

	"github.com/swiggy-private/ads-sos/ads-common/pkg/kafka"
	"github.com/swiggy-private/ads-sos/ads-common/pkg/serviceinit"
)

func main() {
	const (
		KAFKA_CLUSTER_BROKER                = "localhost:9092"
		KAFKA_CLUSTER_TIER                  = 1
		KAFKA_CLUSTER_BROKER_DIALER_TIMEOUT = time.Duration(10)
		KAFKA_TXN_PRIMARY_API_KEY           = "txn_primary_api_key"
		KAFKA_TXN_PRIMARY_API_SECRET        = "txn_primary"
		KAFKA_CONSUMER_CLIENT_ID            = "consumer-client-id"
	)

	topic := "example-topic"
	partition := "0"

	utilityStore, err := serviceinit.InitUtilityStore("KafkaLearning", "dev")
	if err != nil {
		log.Fatalf("error in initializing config,%v", err)
		return
	}

	kafkaConfig := kafka.ProducerConfig{
		ClusterConfig: &kafka.ClusterConfig{
			Brokers:             KAFKA_CLUSTER_BROKER,
			ClusterTier:         KAFKA_CLUSTER_TIER,
			BrokerDialerTimeout: KAFKA_CLUSTER_BROKER_DIALER_TIMEOUT,
			TxnPrimaryAPIKey:    KAFKA_TXN_PRIMARY_API_KEY,
			TxnPrimaryAPISecret: KAFKA_TXN_PRIMARY_API_SECRET,
		},
		MessageTimeout:  500,
		MessageRetries:  3,
		MaxElapsedTime:  1000,
		InitialInterval: 500,
		ClientId:        KAFKA_CONSUMER_CLIENT_ID,
	}

	kafkaProducer, err := kafka.InitProducer(utilityStore, &kafkaConfig)
	if err != nil {
		fmt.Printf("error initiliasing kafka producer for kafka consumer, err: %v", err)
	}

	for {
		message := fmt.Sprintf("Lmao, Event at %s", time.Now().Format(time.RFC3339))
		kafkaProducer.PublishToKafkaWithRetries([]byte(message), topic, partition)
		if err != nil {
			log.Fatal("failed to write messages:", err)
		}
		fmt.Println("Produced:", message)
		time.Sleep(5 * time.Second) // generate an event every 5 seconds
	}
}
