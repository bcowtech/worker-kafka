package kafka

import (
	"fmt"
	"os"
)

var _ NameTransformProc = NameFormatter("")

func NameFormatter(format string) NameTransformProc {
	return func(name string) string {
		return fmt.Sprintf(os.ExpandEnv(format), name)
	}
}
