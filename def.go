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

	Consumer        = kafka.Consumer
	ConsumeContext  = kafka.ConsumeContext
	Forwarder       = kafka.Forwarder
	ForwarderOption = kafka.ForwarderOption
	ForwarderRunner = kafka.ForwarderRunner
	Producer        = kafka.Producer
	ProducerOption  = kafka.ProducerOption

	MessageHandler    = internal.MessageHandler
	NameTransformProc = internal.NameTransformProc
	Worker            = internal.KafkaWorker
)

// func
type (
	KafkaErrorHandleProc = internal.KafkaErrorHandleProc
)
