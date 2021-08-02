package middleware

import (
	"github.com/bcowtech/host"
	"github.com/bcowtech/structproto"
	"github.com/bcowtech/worker-kafka/internal"
)

var _ host.Middleware = new(TopicGatewayMiddleware)

type TopicGatewayMiddleware struct {
	TopicGateway     interface{}
	topicTransformer internal.NameTransformProc
}

func (m *TopicGatewayMiddleware) Init(appCtx *host.AppContext) {
	var (
		kafkaworker = asKafkaWorker(appCtx.Host())
		preparer    = internal.NewKafkaWorkerPreparer(kafkaworker)
	)

	binder := &TopicGatewayBinder{
		router:                           preparer.Router(),
		appContext:                       appCtx,
		configureUnhandledMessageHandler: preparer.RegisterUnhandledMessageHandler,
		topicTransformer:                 m.topicTransformer,
	}

	err := m.performBindTopicGateway(m.TopicGateway, binder)
	if err != nil {
		panic(err)
	}
}

func (m *TopicGatewayMiddleware) TopicTransformer(transformer internal.NameTransformProc) *TopicGatewayMiddleware {
	m.topicTransformer = transformer
	return m
}

func (m *TopicGatewayMiddleware) performBindTopicGateway(target interface{}, binder *TopicGatewayBinder) error {
	prototype, err := structproto.Prototypify(target,
		&structproto.StructProtoResolveOption{
			TagName:     TAG_TOPIC,
			TagResolver: TopicTagResolve,
		},
	)
	if err != nil {
		return err
	}

	return prototype.Bind(binder)
}
