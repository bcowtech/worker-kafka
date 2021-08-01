package middleware

import (
	"github.com/bcowtech/host"
	"github.com/bcowtech/worker-kafka/internal"
)

var _ host.Middleware = new(ErrorHandlerMiddleware)

type ErrorHandlerMiddleware struct {
	Handler internal.KafkaErrorHandleProc
}

func (m *ErrorHandlerMiddleware) Init(appCtx *host.AppContext) {
	var (
		kafkaworker = asKafkaWorker(appCtx.Host())
		preparer    = internal.NewKafkaWorkerPreparer(kafkaworker)
	)

	preparer.RegisterErrorHandler(m.Handler)
}
