package logger

// Logger is the interface of logger instances.
type Logger interface {
	Debugf(tpl string, values ...interface{})
	Infof(tpl string, values ...interface{})
	Warnf(tpl string, values ...interface{})
	Errorf(tpl string, values ...interface{})
	Fatalf(tpl string, values ...interface{})
}
