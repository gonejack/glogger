package basicLogger

import (
	"fmt"
	"github.com/gonejack/glogger/internal/logLevel"
	"log"
	"os"
	"sync"
	"time"
)

type basicLogger struct {
	name   string
	time   string
	stdout *log.Logger
	stderr *log.Logger
	sync.RWMutex
}

func (l *basicLogger) Debugf(tpl string, values ...interface{}) {
	if logLevel.Threshold <= logLevel.DEBUG {
		_ = l.stdout.Output(2, l.format(logLevel.DEBUG, tpl, values...))
	}
}

func (l *basicLogger) Infof(tpl string, values ...interface{}) {
	if logLevel.Threshold <= logLevel.INFO {
		_ = l.stdout.Output(2, l.format(logLevel.INFO, tpl, values...))
	}
}

func (l *basicLogger) Warnf(tpl string, values ...interface{}) {
	if logLevel.Threshold <= logLevel.WARN {
		_ = l.stderr.Output(2, l.format(logLevel.WARN, tpl, values...))
	}
}

func (l *basicLogger) Errorf(tpl string, values ...interface{}) {
	if logLevel.Threshold <= logLevel.ERROR {
		_ = l.stderr.Output(2, l.format(logLevel.ERROR, tpl, values...))
	}
}

func (l *basicLogger) Fatalf(tpl string, values ...interface{}) {
	if logLevel.Threshold <= logLevel.FATAL {
		_ = l.stderr.Output(2, l.format(logLevel.FATAL, tpl, values...))
	}

	os.Exit(-1)
}

func (l *basicLogger) format(level logLevel.LEVEL, tpl string, values ...interface{}) string {
	if len(values) > 0 {
		tpl = fmt.Sprintf(tpl, values...)
	}

	return fmt.Sprintf("[%s] [%s] %s", l.getTime(), level, tpl)
}

func (l *basicLogger) getTime() (time string) {
	l.RLock()
	time = l.time
	l.RUnlock()

	return
}

func New(name string) (logger *basicLogger) {
	logger = &basicLogger{
		name:   name,
		stdout: log.New(os.Stdout, name, 0),
		stderr: log.New(os.Stderr, name, 0),
	}

	go func() {
		for {
			logger.Lock()
			logger.time = time.Now().Format("2006-01-02 15:04:05")
			logger.Unlock()
			time.Sleep(time.Second)
		}
	}()

	return
}
