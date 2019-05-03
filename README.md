# Logger module for go

[![Build Status](https://travis-ci.org/gonejack/glogger.svg?branch=master)](https://travis-ci.org/gonejack/glogger)
[![GoDoc](https://godoc.org/github.com/gonejack/glogger?status.svg)](https://godoc.org/github.com/gonejack/glogger)
[![GitHub license](https://img.shields.io/github/license/gonejack/glogger.svg?color=blue)](LICENSE.md)
```go
var logger = NewLogger("test-logger")

logger.Debugf("Debugf: %s", "Debugf")
logger.Infof("Infof: %s", "Infof")
logger.Warnf("Warnf: %s", "Warnf")
logger.Errorf("Errorf: %s", "Errorf")
```