package internal

import (
	"context"
	"strings"
	"sync"
	"time"

	"github.com/bcowtech/host"
)

var _ host.Host = new(KafkaWorker)

type KafkaWorker struct {
	PollingTimeout time.Duration
	PingTimeout    time.Duration
	ConfigMap      *ConfigMap

	consumer         *Consumer
	bootstrapServers string // e.g: 127.0.0.1:9092,127.0.0.1:9093
	groupID          string
	topics           []string

	dispatcher *KafkaMessageDispatcher

	groupIDTransformer NameTransformProc

	wg          sync.WaitGroup
	mutex       sync.Mutex
	initialized bool
	running     bool
	disposed    bool
}

func (w *KafkaWorker) Start(ctx context.Context) {
	if w.disposed {
		logger.Panic("the Worker has been disposed")
	}
	if !w.initialized {
		logger.Panic("the Worker havn't be initialized yet")
	}
	if w.running {
		return
	}

	var err error
	w.mutex.Lock()
	defer func() {
		if err != nil {
			w.running = false
			w.disposed = true
		}
		w.mutex.Unlock()
	}()

	w.running = true

	c := w.consumer

	logger.Printf("group [%s] listening topics [%s] on address %s\n",
		w.groupID,
		strings.Join(w.topics, ","),
		w.bootstrapServers)

	if len(w.topics) > 0 {
		err = c.Subscribe(w.topics, nil)
		if err != nil {
			logger.Panic(err)
		}
	}
}

func (w *KafkaWorker) Stop(ctx context.Context) error {
	logger.Printf("%% Stopping\n")
	defer func() {
		logger.Printf("%% Stopped\n")
	}()

	w.consumer.Close()
	return nil
}

func (w *KafkaWorker) preInit() {
	w.dispatcher = NewKafkaMessageDispatcher()
	w.groupIDTransformer = NopNameTransformer
}

func (w *KafkaWorker) init() {
	if w.initialized {
		return
	}

	w.mutex.Lock()
	defer func() {
		w.initialized = true
		w.mutex.Unlock()
	}()

	w.configTopics()
	w.configBootstrapServers()
	w.configGroupID()
	w.configConsumer()
}

func (w *KafkaWorker) configTopics() {
	w.topics = w.dispatcher.Topics()
}

func (w *KafkaWorker) configBootstrapServers() {
	v, _ := w.ConfigMap.Get(KAFKA_CONF_BOOTSTRAP_SERVERS, nil)
	if v == nil {
		logger.Panicf("missing kafka config '%s'", KAFKA_CONF_BOOTSTRAP_SERVERS)
	}
	w.bootstrapServers = v.(string)
}

func (w *KafkaWorker) configGroupID() {
	v, _ := w.ConfigMap.Get(KAFKA_CONF_GROUP_ID, nil)
	if v == nil {
		logger.Panicf("missing kafka config '%s'", KAFKA_CONF_GROUP_ID)
	}
	if groupID, ok := v.(string); ok {
		// transform
		realGroupID := w.groupIDTransformer(groupID)
		// update ConfigMap
		err := w.ConfigMap.SetKey(KAFKA_CONF_GROUP_ID, realGroupID)
		if err != nil {
			panic(err)
		}
		// export
		w.groupID = realGroupID
	}
}

func (w *KafkaWorker) configConsumer() {
	instance := &Consumer{
		ConfigMap:               w.ConfigMap,
		PollingTimeout:          w.PollingTimeout,
		PingTimeout:             w.PingTimeout,
		ErrorHandler:            w.dispatcher.ProcessKafkaError,
		MessageHandler:          w.dispatcher.ProcessMessage,
		UnhandledMessageHandler: w.dispatcher.ProcessUnhandledMessage,
	}

	w.consumer = instance
}
