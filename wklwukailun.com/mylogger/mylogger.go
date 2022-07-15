package mylogger

import (
	"errors"
	"os"
	"path"
	"runtime"
	"strings"
)

/*
1.支持向不同的地方输出日志
2.日志级别
debug trace info warning Error fatal
3.日志需要开关控制,控制日志输出级别
4.完整的日志记录需要有时间、行号、文件名、日志级别、日志信息
5.日志文件切割
*/
type LogLevel uint16
type ConsoleLogger struct {
	Level LogLevel
}
type FileLogger struct {
	Level       LogLevel
	filePath    string
	fileName    string
	maxFileSize int64
	fileObj     *os.File
	errfileObj  *os.File
	logChan     chan *logMsg
}
type logMsg struct {
	Level     LogLevel
	funcName  string
	fileName  string
	lineNo    int
	msg       string
	timestamp string
}

const (
	DEBUG LogLevel = iota
	TRACE
	INFO
	WARNING
	ERROR
	FATAL
)

func parseLogLevel(s string) (LogLevel, error) {
	switch s = strings.ToLower(s); s {
	case "debug":
		return DEBUG, nil
	case "trace":
		return TRACE, nil
	case "info":
		return INFO, nil
	case "warning":
		return WARNING, nil
	case "error":
		return ERROR, nil
	case "fatal":
		return FATAL, nil
	default:
		return DEBUG, errors.New("无效的级别")
	}
}
func getLogString(lv LogLevel) string {
	switch lv {
	case DEBUG:
		return "DEBUG"
	case TRACE:
		return "TRACE"
	case INFO:
		return "INFO"
	case WARNING:
		return "WARNING"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	default:
		return ""
	}
}
func getInfo(n int) (funcName, fileName string, lineNo int) {
	pc, file, lineNo, ok := runtime.Caller(n)
	if !ok {
		return
	}
	funcName = runtime.FuncForPC(pc).Name()
	fileName = path.Base(file)
	return
}
