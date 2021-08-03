package internal

import "github.com/confluentinc/confluent-kafka-go/kafka"

type KafkaMessageDispatcher struct {
	kafkaErrorHandler       KafkaErrorHandleProc
	messageHandler          MessageHandleProc
	unhandledMessageHandler MessageHandler
	router                  Router
}

func NewKafkaMessageDispatcher() *KafkaMessageDispatcher {
	return &KafkaMessageDispatcher{
		router: make(Router),
	}
}

func (d *KafkaMessageDispatcher) Topics() []string {
	var (
		router = d.router
	)

	if router != nil {
		keys := make([]string, 0, len(router))
		for k := range router {
			keys = append(keys, k)
		}
		return keys
	}
	return nil
}

func (d *KafkaMessageDispatcher) ProcessMessage(worker *ConsumeWorker, message *Message) {

	// TODO: handle error
	// defer func() {
	// 	err := recover()
	// 	if err != nil {
	// 		d.errorHandler(err)
	// 	}
	// }()

	var (
		topic string = *message.TopicPartition.Topic
	)

	handler := d.router.Get(topic)
	if handler != nil {
		handler.ProcessMessage(worker, message)
	} else {
		worker.ForwardUnhandledMessage(message)
	}
}

func (d *KafkaMessageDispatcher) ProcessUnhandledMessage(worker *ConsumeWorker, message *Message) {
	if d.unhandledMessageHandler != nil {
		d.unhandledMessageHandler.ProcessMessage(worker, message)
	}
}

func (d *KafkaMessageDispatcher) ProcessKafkaError(err kafka.Error) (disposed bool) {
	if d.kafkaErrorHandler != nil {
		return d.kafkaErrorHandler(err)
	}
	return false
}
