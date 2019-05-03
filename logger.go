/*
Package glogger is a basic logger library for simple go projects.
*/
package glogger

import "github.com/gonejack/glogger/internal/basicLogger"

// NewLogger creates a new logger instance match Logger interface
func NewLogger(name string) Logger {
	if name != "" {
		name = "[" + name + "] "
	}

	return basicLogger.New(name)
}
