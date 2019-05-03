# Logger module for go


```$go
var logger = NewLogger("test-logger")

logger.Debugf("Debugf: %s", "Debugf")
logger.Infof("Infof: %s", "Infof")
logger.Warnf("Warnf: %s", "Warnf")
logger.Errorf("Errorf: %s", "Errorf")
```