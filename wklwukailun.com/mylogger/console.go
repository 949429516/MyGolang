package mylogger

import (
	"fmt"
	"time"
)

//向终端输出日志内容

func NewConsoleLog(levelStr string) ConsoleLogger {
	level, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	return ConsoleLogger{
		Level: level,
	}
}

func (l ConsoleLogger) consolelog(lv LogLevel, format string, a ...interface{}) {
	if lv >= l.Level {
		msg := fmt.Sprintf(format, a...)
		now := time.Now()
		funcName, fileName, lineNo := getInfo(3)
		fmt.Printf("[%s] [%s] [%s %s %d] %s\n", now.Format("2006-01-02 15:04:05"), getLogString(lv), funcName, fileName, lineNo, msg)
	}
}

func (l ConsoleLogger) Debug(format string, a ...interface{}) {
	l.consolelog(DEBUG, format, a...)
}
func (l ConsoleLogger) Info(format string, a ...interface{}) {
	l.consolelog(INFO, format, a...)
}
func (l ConsoleLogger) Trace(format string, a ...interface{}) {
	l.consolelog(TRACE, format, a...)
}
func (l ConsoleLogger) Warning(format string, a ...interface{}) {
	l.consolelog(WARNING, format, a...)
}
func (l ConsoleLogger) Error(format string, a ...interface{}) {
	l.consolelog(ERROR, format, a...)
}
func (l ConsoleLogger) Fatal(format string, a ...interface{}) {
	l.consolelog(FATAL, format, a...)
}
