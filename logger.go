package metadata

import "github.com/phonkee/go-logger"

var (
	// global debug option
	debug = false
)

var (
	loggerInfo    = logger.Info("metadata")
	loggerError   = logger.Error("metadata")
	loggerDebug   = logger.Debug("metadata")
	loggerWarning = logger.Warning("metadata")
)

