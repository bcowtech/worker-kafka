package middleware

import (
	"fmt"
	"reflect"

	"github.com/bcowtech/host"
	"github.com/bcowtech/structproto"
	"github.com/bcowtech/structproto/tagresolver"
	"github.com/bcowtech/structproto/util/reflectutil"
	"github.com/bcowtech/worker-kafka/internal"
)

var _ structproto.StructBinder = new(TopicGatewayBinder)

type TopicGatewayBinder struct {
	appContext                       *host.AppContext
	router                           internal.Router
	configureUnhandledMessageHandler ConfigureUnhandledMessageHandleProc
	topicTransformer                 internal.NameTransformProc
}

func (b *TopicGatewayBinder) Init(context *structproto.StructProtoContext) error {
	return nil
}

func (b *TopicGatewayBinder) Bind(field structproto.FieldInfo, rv reflect.Value) error {
	if !rv.IsValid() {
		return fmt.Errorf("specifiec argument 'rv' is invalid")
	}

	// assign zero if rv is nil
	rvMessageHandler := reflectutil.AssignZero(rv)
	binder := &MessageHandlerBinder{
		messageHandlerType: rv.Type().Name(),
		components: map[string]reflect.Value{
			host.APP_CONFIG_FIELD:           b.appContext.Config(),
			host.APP_SERVICE_PROVIDER_FIELD: b.appContext.ServiceProvider(),
		},
	}
	err := b.preformBindMessageHandler(rvMessageHandler, binder)
	if err != nil {
		return err
	}

	// register MessageHandlers
	return b.registerRoute(field.Name(), rvMessageHandler)
}

func (b *TopicGatewayBinder) Deinit(context *structproto.StructProtoContext) error {
	return nil
}

func (b *TopicGatewayBinder) preformBindMessageHandler(target reflect.Value, binder *MessageHandlerBinder) error {
	prototype, err := structproto.Prototypify(target,
		&structproto.StructProtoResolveOption{
			TagResolver: tagresolver.NoneTagResolver,
		})
	if err != nil {
		return err
	}

	return prototype.Bind(binder)
}

func (b *TopicGatewayBinder) registerRoute(topic string, rv reflect.Value) error {
	// register MessageHandlers
	if isMessageHandler(rv) {
		handler := asMessageHandler(rv)
		if handler != nil {
			if topic == UNHANDLED_MESSAGE_HANDLER_TOPIC_SYMBOL {
				b.configureUnhandledMessageHandler(handler)
			} else {
				realTopic := b.transformTopic(topic)
				b.router.Add(realTopic, handler)
			}
		}
	}
	return nil
}

func (b *TopicGatewayBinder) transformTopic(topic string) string {
	if b.topicTransformer != nil {
		return b.topicTransformer(topic)
	}
	return topic
}
