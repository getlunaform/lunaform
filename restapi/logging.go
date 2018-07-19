package restapi

import "github.com/go-openapi/runtime/logger"

// Debug when true turns on verbose logging
var Debug = logger.DebugEnabled()
var Logger logger.Logger = logger.StandardLogger{}

func DebugLog(format string, args ...interface{}) {
	if Debug {
		Logger.Printf(format, args...)
	}
}
