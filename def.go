package kafka

import (
	kafka "github.com/bcowtech/lib-kafka"
	"github.com/bcowtech/worker-kafka/internal"
)

const (
	KAFKA_CONF_BOOTSTRAP_SERVERS = kafka.KAFKA_CONF_BOOTSTRAP_SERVERS
	KAFKA_CONF_GROUP_ID          = "group.id"

	PartitionAny = kafka.PartitionAny
)

// interface & struct
type (
	ConsumerGroupMetadata = kafka.ConsumerGroupMetadata
	ConfigMap             = kafka.ConfigMap
	Error                 = kafka.Error
	ErrorCode             = kafka.ErrorCode
	Event                 = kafka.Event
	Message               = kafka.Message
	Offset                = kafka.Offset
	RebalanceCb           = kafka.RebalanceCb
	TopicPartition        = kafka.TopicPartition

	Forwarder       = kafka.Forwarder
	ForwarderOption = kafka.ForwarderOption
	ForwarderRunner = kafka.ForwarderRunner
	WorkerContext   = kafka.ConsumeContext

	MessageHandler = internal.MessageHandler
	Worker         = internal.KafkaWorker
)

// func
type (
	NameTransformProc    = internal.NameTransformProc
	KafkaErrorHandleProc = internal.KafkaErrorHandleProc
)
