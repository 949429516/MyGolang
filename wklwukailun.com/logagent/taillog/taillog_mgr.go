package taillog

import "wklwukailun.com/logagent/etcd"

type taillogMgr struct {
	logEntry []*etcd.LogEntry
	//tskMap   map[string]*TailTask
}

var tskMgr *taillogMgr

func Init(logEntryConf []*etcd.LogEntry) {
	tskMgr = &taillogMgr{
		logEntry: logEntryConf, // 把当前日志收集项保存起来
	}
	for _, logEntry := range logEntryConf {
		// logEntry : *etcd.LogEntry   logEntry.Path要收集日志的路径
		NewTailTask(logEntry.Path, logEntry.Topic)
	}
}
