package common

import (
	"github.com/digitalbitsorg/go/support/log"
)

const StellarAmountPrecision = 7

func CreateLogger(serviceName string) *log.Entry {
	return log.DefaultLogger.WithField("service", serviceName)
}
