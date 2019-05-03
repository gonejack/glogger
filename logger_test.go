package glogger

import (
	"testing"
	"time"
)

func TestNewLogger(t *testing.T) {
	var logger = NewLogger("test-logger")

	logger.Debugf("Debugf: %s", "Debugf")
	time.Sleep(time.Second)

	logger.Infof("Infof: %s", "Infof")
	time.Sleep(time.Second)

	logger.Warnf("Warnf: %s", "Warnf")
	time.Sleep(time.Second)

	logger.Errorf("Errorf: %s", "Errorf")
	time.Sleep(time.Second)

	//logger.Fatalf("Fatalf: %s", "Fatalf")

	t.Logf("done")
}
