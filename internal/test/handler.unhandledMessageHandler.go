package test

import (
	"fmt"

	kafka "github.com/bcowtech/worker-kafka"
)

type UnhandledMessageHandler struct {
	ServiceProvider *ServiceProvider
}

func (h *UnhandledMessageHandler) Init() {
	fmt.Println("UnhandledMessageHandler.Init()")
}

func (h *UnhandledMessageHandler) ProcessMessage(ctx *kafka.WorkerContext, message *kafka.Message) {
	fmt.Println("UnhandledMessageHandler.ProcessMessage()")
}
