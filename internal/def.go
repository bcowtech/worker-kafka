package internal

import (
	"log"
	"os"
	"reflect"

	kafka "github.com/bcowtech/lib-kafka"
)

const (
	KAFKA_CONF_BOOTSTRAP_SERVERS = "bootstrap.servers"
	KAFKA_CONF_GROUP_ID          = "group.id"

	LOGGER_PREFIX = "[bcowtech/worker-kafka] "
)

var (
	KafkaWorkerServiceInstance = new(KafkaWorkerService)

	typeOfHost = reflect.TypeOf(KafkaWorker{})

	logger *log.Logger = log.New(os.Stdout, LOGGER_PREFIX, log.LstdFlags|log.Lmsgprefix)
)

// interface & struct
type (
	Consumer       = kafka.Consumer
	ConsumeContext = kafka.ConsumeContext
	ConfigMap      = kafka.ConfigMap
	Message        = kafka.Message
	TopicPartition = kafka.TopicPartition
	Offset         = kafka.Offset
	Error          = kafka.Error
	ErrorCode      = kafka.ErrorCode
	RebalanceCb    = kafka.RebalanceCb

	MessageHandler interface {
		ProcessMessage(ctx *ConsumeContext, message *Message)
	}
)

// func
type (
	MessageHandleProc    = kafka.MessageHandleProc
	KafkaErrorHandleProc = kafka.ErrorHandleProc

	NameTransformProc func(name string) string
)
