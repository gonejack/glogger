# Logger module for go

[![Build Status](https://travis-ci.org/gonejack/logger.svg?branch=master)](https://travis-ci.org/gonejack/logger)
[![GoDoc](https://godoc.org/github.com/gonejack/logger?status.svg)](https://godoc.org/github.com/gonejack/logger)
[![GitHub license](https://img.shields.io/github/license/gonejack/logger.svg?color=blue)](LICENSE.md)
```go
var logger = NewLogger("test-logger")

logger.Debugf("Debugf: %s", "Debugf")
logger.Infof("Infof: %s", "Infof")
logger.Warnf("Warnf: %s", "Warnf")
logger.Errorf("Errorf: %s", "Errorf")
```