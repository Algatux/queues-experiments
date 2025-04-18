package internal

import (
	"github.com/pho3b/tiny-logger/logs"
	"github.com/pho3b/tiny-logger/logs/log_level"
)

var Logger *logs.Logger

func init() {
	Logger = logs.NewLogger().
		SetLogLvl(log_level.DebugLvlName).
		EnableColors(true).
		AddTime(true)
}
