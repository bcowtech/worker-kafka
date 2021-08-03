package kafka

import kafka "github.com/bcowtech/lib-kafka"

func NewForwarder(opt *ForwarderOption) (*kafka.Forwarder, error) {
	return kafka.NewForwarder(opt)
}
