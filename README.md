# Logger module for go

[![Build Status](https://travis-ci.org/gonejack/logger.svg?branch=master)](https://travis-ci.org/gonejack/logger)

```go
var logger = NewLogger("test-logger")

logger.Debugf("Debugf: %s", "Debugf")
logger.Infof("Infof: %s", "Infof")
logger.Warnf("Warnf: %s", "Warnf")
logger.Errorf("Errorf: %s", "Errorf")
```