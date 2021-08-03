package kafka

import (
	"fmt"
	"os"

	kafka "github.com/bcowtech/lib-kafka"
)

var _ NameTransformProc = NameFormatter("")

func NameFormatter(format string) NameTransformProc {
	return func(name string) string {
		return fmt.Sprintf(os.ExpandEnv(format), name)
	}
}

func NewForwarder(opt *ForwarderOption) (*kafka.Forwarder, error) {
	return kafka.NewForwarder(opt)
}
