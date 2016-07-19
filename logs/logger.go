package logs

import (
	"github.com/astaxie/beego/logs"
	"io/ioutil"
)

var (
	DEBUG  bool
	logger *logs.BeeLogger
)

func init() {
	logger = logs.NewLogger(1)
	logger.EnableFuncCallDepth(true)
	logger.SetLogFuncCallDepth(2)
}

func Init(file string) error {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	return InitByString(string(content))
}

func InitByString(content string) error {
	logger = logs.NewLogger(1)
	logger.EnableFuncCallDepth(true)
	logger.SetLogFuncCallDepth(2)

	logger.SetLogger("file", content)

	if DEBUG {
		logger.SetLevel(logs.LevelDebug)
	} else {
		logger.SetLevel(logs.LevelInfo)
		logger.DelLogger("console")
	}
	logger.Info("startint...")

	return nil
}

func Async() *logs.BeeLogger {
	return logger.Async()
}

// SetLogger provides a given logger adapter into BeeLogger with config string.
// config need to be correct JSON as string: {"interval":360}.
func SetLogger(adapterName string, config string) error {
	return logger.SetLogger(adapterName, config)
}

// DelLogger remove a logger adapter in BeeLogger.
func DelLogger(adapterName string) error {
	return logger.DelLogger(adapterName)
}

// SetLevel Set log message level.
// If message level (such as LevelDebug) is higher than logger level (such as LevelWarning),
// log providers will not even be sent the message.
func SetLevel(l int) {
	logger.SetLevel(l)
}

// SetLogFuncCallDepth set log funcCallDepth
func SetLogFuncCallDepth(d int) {
	logger.SetLogFuncCallDepth(d)
}

// GetLogFuncCallDepth return log funcCallDepth for wrapper
func GetLogFuncCallDepth() int {
	return logger.GetLogFuncCallDepth()
}

// EnableFuncCallDepth enable log funcCallDepth
func EnableFuncCallDepth(b bool) {
	logger.EnableFuncCallDepth(b)
}

// Emergency Log EMERGENCY level message.
func Emergency(format string, v ...interface{}) {
	logger.Emergency(format, v)
}

// Alert Log ALERT level message.
func Alert(format string, v ...interface{}) {
	logger.Alert(format, v)
}

// Critical Log CRITICAL level message.
func Critical(format string, v ...interface{}) {
	logger.Critical(format, v)
}

// Error Log ERROR level message.
func Error(format string, v ...interface{}) {
	logger.Error(format, v)
}

// Warning Log WARNING level message.
func Warning(format string, v ...interface{}) {
	logger.Warning(format, v)
}

// Notice Log NOTICE level message.
func Notice(format string, v ...interface{}) {
	logger.Notice(format, v)
}

// Informational Log INFORMATIONAL level message.
func Informational(format string, v ...interface{}) {
	logger.Informational(format, v)
}

// Debug Log DEBUG level message.
func Debug(format string, v ...interface{}) {
	logger.Debug(format, v)
}

// Warn Log WARN level message.
// compatibility alias for Warning()
func Warn(format string, v ...interface{}) {
	logger.Warn(format, v)
}

// Info Log INFO level message.
// compatibility alias for Informational()
func Info(format string, v ...interface{}) {
	logger.Info(format, v)
}

// Trace Log TRACE level message.
// compatibility alias for Debug()
func Trace(format string, v ...interface{}) {
	logger.Trace(format, v)
}

// Flush flush all chan data.
func Flush() {
	logger.Flush()
}

// Close close logger, flush all chan data and destroy all adapters in BeeLogger.
func Close() {
	logger.Close()
}

// Reset close all outputs, and set bl.outputs to nil
func Reset() {
	logger.Reset()
}
