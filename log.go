package logger

import "github.com/gonejack/logger/internal/basicLogger"

func NewLogger(name string) Logger {
	if name != "" {
		name = "[" + name + "] "
	}

	return basicLogger.New(name)
}

