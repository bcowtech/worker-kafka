package kafka

import (
	"github.com/bcowtech/host"
	"github.com/bcowtech/worker-kafka/internal"
)

func Startup(app interface{}) *host.Starter {
	var (
		starter = host.Startup(app)
	)

	host.RegisterHostService(starter, internal.KafkaWorkerServiceInstance)

	return starter
}
