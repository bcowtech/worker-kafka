package test

import (
	"fmt"
	"log"

	kafka "github.com/bcowtech/worker-kafka"
)

var _ kafka.MessageHandler = new(MyTopicMessageHandler)

type MyTopicMessageHandler struct {
	ServiceProvider *ServiceProvider
}

func (h *MyTopicMessageHandler) Init() {
	fmt.Println("MyTopicMessageHandler.Init()")
}

func (h *MyTopicMessageHandler) ProcessMessage(ctx *kafka.WorkerContext, message *kafka.Message) {
	log.Printf("Message on %s: %s: %s\n", message.TopicPartition, string(message.Key), string(message.Value))
}
