/*
Package logger is a basic logger library for simple go projects.
*/
package glogger

import "github.com/gonejack/logger/internal/basicLogger"

// NewLogger creates a new logger instance match Logger interface
func NewLogger(name string) Logger {
	if name != "" {
		name = "[" + name + "] "
	}

	return basicLogger.New(name)
}
