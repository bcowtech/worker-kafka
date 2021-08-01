package internal

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"
)

func TestKafkaWorker(t *testing.T) {
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
