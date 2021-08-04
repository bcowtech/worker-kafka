package internal

type KafkaWorkerPreparer struct {
	worker *KafkaWorker
}

func NewKafkaWorkerPreparer(worker *KafkaWorker) *KafkaWorkerPreparer {
	return &KafkaWorkerPreparer{
		worker: worker,
	}
}

func (p *KafkaWorkerPreparer) RegisterErrorHandler(handler KafkaErrorHandleProc) {
	p.worker.dispatcher.kafkaErrorHandler = handler
}

func (p *KafkaWorkerPreparer) RegisterGroupIDTransformer(transformer NameTransformProc) {
	p.worker.groupIDTransformer = transformer
}

func (p *KafkaWorkerPreparer) RegisterUnhandledMessageHandler(handler MessageHandler) {
	p.worker.dispatcher.unhandledMessageHandler = handler
}

func (p *KafkaWorkerPreparer) Router() Router {
	return p.worker.dispatcher.router
}
