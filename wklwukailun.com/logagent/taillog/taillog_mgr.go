package taillog

import (
	"fmt"
	"time"

	"wklwukailun.com/logagent/etcd"
)

// tailTask管理者
type taillogMgr struct {
	logEntry    []*etcd.LogEntry
	tskMap      map[string]*TailTask
	newConfChan chan []*etcd.LogEntry
}

var tskMgr *taillogMgr

func Init(logEntryConf []*etcd.LogEntry) {
	tskMgr = &taillogMgr{
		logEntry:    logEntryConf, // 把当前日志收集项保存起来
		tskMap:      make(map[string]*TailTask, 16),
		newConfChan: make(chan []*etcd.LogEntry), //无缓冲区的通道
	}
	for _, logEntry := range logEntryConf {
		// logEntry : *etcd.LogEntry   logEntry.Path要收集日志的路径
		NewTailTask(logEntry.Path, logEntry.Topic)
	}
	go tskMgr.run()
}

// 监听自己的通道,有新的配置过来就做处理1.配置新增 2.配置删除 3.配置变更
func (t *taillogMgr) run() {
	for {
		select {
		case newConf := <-t.newConfChan:
			fmt.Println("新的配置来了", newConf)

		default:
			time.Sleep(time.Second)
		}
	}
}

// 向外暴露一个函数向tskMgr的newConfChan
func NewConfChan() chan<- []*etcd.LogEntry {
	return tskMgr.newConfChan
}
