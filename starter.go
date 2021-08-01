package kafka

import (
	"github.com/bcowtech/host"
	"github.com/bcowtech/worker-kafka/internal"
)

func Startup(app interface{}, middlewares ...host.Middleware) *host.Starter {
	var (
		starter = host.Startup(app, middlewares...)
	)

	host.RegisterHostService(starter, internal.KafkaWorkerServiceInstance)

	return starter
}
