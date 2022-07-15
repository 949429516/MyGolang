package mylogger

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"time"
)

func NewFileLog(levelStr, fp, fn string, maxFileSize int64) *FileLogger {
	level, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	fullFileName := filepath.Join(fp, fn)
	fileobj, err := os.OpenFile(fullFileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	fileerrobj, err := os.OpenFile(fullFileName+".err", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	return &FileLogger{
		Level:       level,
		filePath:    fp,
		fileName:    fn,
		maxFileSize: maxFileSize,
		fileObj:     fileobj,
		errfileObj:  fileerrobj,
		logChan:     make(chan *logMsg, 50000),
	}
}
func (l *FileLogger) checkSize(file *os.File) bool {
	fileObj, err := file.Stat()
	if err != nil {
		return false
	}
	return fileObj.Size() >= l.maxFileSize
}
func (l *FileLogger) splitFile(file *os.File) (*os.File, error) {
	nowStr := time.Now().Format("20060102150405000000")
	logName := path.Join(l.filePath, file.Name())
	newLogName := fmt.Sprintf("%s.bak%s", logName, nowStr)
	os.Rename(logName, newLogName)
	file.Close()
	fileobj, err := os.OpenFile(logName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	return fileobj, nil
}
func (l *FileLogger) writeLogBackground() {
	for {
		if l.checkSize(l.fileObj) {
			newFile, err := l.splitFile(l.fileObj)
			if err != nil {
				return
			}
			l.fileObj = newFile
		}
		select {
		case logTmp := <-l.logChan:
			fmt.Fprintf(l.fileObj, "[%s] [%s] [%s %s %d] %s\n", logTmp.timestamp, getLogString(logTmp.Level), logTmp.funcName, logTmp.fileName, logTmp.lineNo, logTmp.msg)
			if logTmp.Level >= ERROR {
				if l.checkSize(l.errfileObj) {
					newFile, err := l.splitFile(l.errfileObj)
					if err != nil {
						return
					}
					l.errfileObj = newFile
				}
				fmt.Fprintf(l.errfileObj, "[%s] [%s] [%s %s %d] %s\n", logTmp.timestamp, getLogString(logTmp.Level), logTmp.funcName, logTmp.fileName, logTmp.lineNo, logTmp.msg)
			}
		default:
			//取不到日志休息500ms
			time.Sleep(time.Millisecond * 500)
		}
	}
}

func (l *FileLogger) filelog(lv LogLevel, format string, a ...interface{}) {
	go l.writeLogBackground()
	if lv >= l.Level {
		msg := fmt.Sprintf(format, a...)
		now := time.Now()
		funcName, fileName, lineNo := getInfo(3)
		//日志发送到通道中
		logTmp := &logMsg{
			Level:     lv,
			funcName:  funcName,
			fileName:  fileName,
			lineNo:    lineNo,
			msg:       msg,
			timestamp: now.Format("2006-01-02 15:04:05"),
		}
		select {
		case l.logChan <- logTmp:
		default:
			//考虑极端情况,管道写满后会阻塞，为保证业务代码正常丢弃日志
		}
	}
}
func (l *FileLogger) Debug(format string, a ...interface{}) {
	l.filelog(DEBUG, format, a...)
}
func (l *FileLogger) Info(format string, a ...interface{}) {
	l.filelog(INFO, format, a...)
}
func (l *FileLogger) Trace(format string, a ...interface{}) {
	l.filelog(TRACE, format, a...)
}
func (l *FileLogger) Warning(format string, a ...interface{}) {
	l.filelog(WARNING, format, a...)
}
func (l *FileLogger) Error(format string, a ...interface{}) {
	l.filelog(ERROR, format, a...)
}
func (l *FileLogger) Fatal(format string, a ...interface{}) {
	l.filelog(FATAL, format, a...)
}
