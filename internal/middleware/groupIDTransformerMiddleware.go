package middleware

import (
	"github.com/bcowtech/host"
	"github.com/bcowtech/worker-kafka/internal"
)

var _ host.Middleware = new(GroupIDTransformerMiddleware)

type GroupIDTransformerMiddleware struct {
	Transformer internal.NameTransformProc
}

func (m *GroupIDTransformerMiddleware) Init(appCtx *host.AppContext) {
	var (
		kafkaworker = asKafkaWorker(appCtx.Host())
		preparer    = internal.NewKafkaWorkerPreparer(kafkaworker)
	)

	preparer.RegisterGroupIDTransformer(m.Transformer)
}
