package internal

import (
	"reflect"

	"github.com/bcowtech/host"
)

var _ host.HostService = new(KafkaWorkerService)

type KafkaWorkerService struct{}

func (s *KafkaWorkerService) Init(h host.Host, ctx *host.AppContext) {
	if v, ok := h.(*KafkaWorker); ok {
		v.preInit()
	}
}

func (s *KafkaWorkerService) InitComplete(h host.Host, ctx *host.AppContext) {
	if v, ok := h.(*KafkaWorker); ok {
		v.init()
	}
}

func (s *KafkaWorkerService) GetHostType() reflect.Type {
	return typeOfHost
}
