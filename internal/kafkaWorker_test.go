package internal

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	kafka "github.com/bcowtech/lib-kafka"
)

func TestKafkaWorker(t *testing.T) {
	err := setupTest()
	if err != nil {
		t.Fatal(err)
	}

	worker := &KafkaWorker{
		PollingTimeout: 30 * time.Millisecond,
		PingTimeout:    3 * time.Second,
		ConfigMap: &ConfigMap{
			"bootstrap.servers": os.Getenv("KAFKA_BOOTSTRAP_SERVERS"),
			"group.id":          "gotest",
			"auto.offset.reset": "earliest",
		},
	}

	worker.preInit()
	{
		worker.dispatcher.router.Add("myTopic", new(mockMessageHandler))
	}
	worker.init()

	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	worker.Start(ctx)

	select {
	case <-ctx.Done():
		worker.Stop(context.Background())
		break
	}
}

type mockMessageHandler struct{}

func (h *mockMessageHandler) ProcessMessage(worker *ConsumeWorker, message *Message) {
	fmt.Printf("Message on %s: %s: %s\n", message.TopicPartition, string(message.Key), string(message.Value))
}

func setupTest() error {
	p, err := kafka.NewProducer(&kafka.ProducerOption{
		FlushTimeout: 3 * time.Second,
		PingTimeout:  3 * time.Second,
		ConfigMap: &kafka.ConfigMap{
			"client.id":         "gotest",
			"bootstrap.servers": os.Getenv("KAFKA_BOOTSTRAP_SERVERS"),
		},
	})
	if err != nil {
		return err
	}

	topic := "myTopic"
	for _, word := range []string{"Welcome", "to", "the", "Confluent", "Kafka", "Golang", "client"} {
		p.WriteMessage(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          []byte(word),
		}, nil)
	}
	p.Close()
	return nil
}
