package test

import (
	"github.com/digitalbitsorg/go/services/frontier/internal/log"
	"github.com/sirupsen/logrus"
)

var testLogger *log.Entry

func init() {
	testLogger, _ = log.New()
	testLogger.Entry.Logger.Formatter.(*logrus.TextFormatter).DisableColors = true
	testLogger.Entry.Logger.Level = logrus.DebugLevel
}
