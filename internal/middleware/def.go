package middleware

import (
	"reflect"

	"github.com/bcowtech/worker-kafka/internal"
)

const (
	UNHANDLED_MESSAGE_HANDLER_TOPIC_SYMBOL string = "?"
)

var (
	typeOfHost           = reflect.TypeOf(internal.KafkaWorker{})
	typeOfMessageHandler = reflect.TypeOf((*internal.MessageHandler)(nil)).Elem()

	TAG_TOPIC = "topic"
)

type (
	ConfigureUnhandledMessageHandleProc func(handler internal.MessageHandler)
)
