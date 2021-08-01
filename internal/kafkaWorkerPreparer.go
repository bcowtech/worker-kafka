package internal

type KafkaWorkerPreparer struct {
	subject *KafkaWorker
}

func NewKafkaWorkerPreparer(subject *KafkaWorker) *KafkaWorkerPreparer {
	return &KafkaWorkerPreparer{
		subject: subject,
	}
}

func (p *KafkaWorkerPreparer) RegisterErrorHandler(handler KafkaErrorHandleProc) {
	p.subject.dispatcher.kafkaErrorHandler = handler
}

func (p *KafkaWorkerPreparer) RegisterGroupIDTransformer(transformer NameTransformProc) {
	p.subject.groupIDTransformer = transformer
}

func (p *KafkaWorkerPreparer) RegisterUnhandledMessageHandler(handler MessageHandler) {
	p.subject.dispatcher.unhandledMessageHandler = handler
}

func (p *KafkaWorkerPreparer) Router() Router {
	return p.subject.dispatcher.router
}
