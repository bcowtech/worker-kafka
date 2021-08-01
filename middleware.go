package kafka

import (
	"github.com/bcowtech/host"
	"github.com/bcowtech/worker-kafka/internal/middleware"
)

func UseErrorHandler(handler KafkaErrorHandleProc) host.Middleware {
	if handler == nil {
		panic("argument 'handler' cannot be nil")
	}

	return &middleware.ErrorHandlerMiddleware{
		Handler: handler,
	}
}

func UseGroupIDTransformer(transformer NameTransformProc) host.Middleware {
	if transformer == nil {
		panic("argument 'transformer' cannot be nil")
	}

	return &middleware.GroupIDTransformerMiddleware{
		Transformer: transformer,
	}
}

func UseTopicGateway(topicGateway interface{}) *middleware.TopicGatewayMiddleware {
	if topicGateway == nil {
		panic("argument 'topicGateway' cannot be nil")
	}

	return &middleware.TopicGatewayMiddleware{
		TopicGateway: topicGateway,
	}
}
