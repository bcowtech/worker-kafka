package main

import (
	"fmt"
	"os"
	"time"

	kafka "github.com/bcowtech/lib-kafka"
)

func main() {
	p, err := kafka.NewProducer(&kafka.ProducerOption{
		FlushTimeout: 3 * time.Second,
		PingTimeout:  3 * time.Second,
		ConfigMap: &kafka.ConfigMap{
			"client.id":         "gotest",
			"bootstrap.servers": os.Getenv("KAFKA_BOOTSTRAP_SERVERS"),
		},
	})
	if err != nil {
		panic(err)
	}

	topic := "myTopic"
	realTopic := fmt.Sprintf(os.ExpandEnv("%s-${KAFKA_TOPIC_SUFFIX}"), topic)
	for _, word := range []string{"Welcome", "to", "the", "Confluent", "Kafka", "Golang", "client"} {
		p.WriteMessage(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &realTopic, Partition: kafka.PartitionAny},
			Value:          []byte(word),
		}, nil)
	}
	p.Close()
}
