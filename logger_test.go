package logger

import (
	"testing"
)

func TestNewLogger(t *testing.T) {
	var logger = NewLogger("test-logger")

	logger.Debugf("Debugf: %s", "Debugf")
	logger.Infof("Infof: %s", "Infof")
	logger.Warnf("Warnf: %s", "Warnf")
	logger.Errorf("Errorf: %s", "Errorf")
	//logger.Fatalf("Fatalf: %s", "Fatalf")

	t.Logf("done")
}
