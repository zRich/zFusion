package flogging

import (
	flogging "github.com/hyperledger/fabric/common/flogging"
)

//temporarily brow fabric log module
type Logger = flogging.FabricLogger

func MustGetLogger(loggerName string) *Logger {
	return flogging.MustGetLogger(loggerName)
}
