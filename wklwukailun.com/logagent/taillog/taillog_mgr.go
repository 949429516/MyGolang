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
		// 有一个配置项启动一个tailtask
		tailtask := NewTailTask(logEntry.Path, logEntry.Topic)
		// 将tailtask存入tskmap
		mk := fmt.Sprintf("%s_%s", logEntry.Path, logEntry.Topic)
		tskMgr.tskMap[mk] = tailtask
	}
	// 后台执行哨兵
	go tskMgr.run()
}

// 监听自己的通道,有新的配置过来就做处理1.配置新增 2.配置删除 3.配置变更
func (t *taillogMgr) run() {
	for {
		select {
		case newConf := <-t.newConfChan:
			for _, conf := range newConf {
				// 判断conf在tskMap里面有没有
				mk := fmt.Sprintf("%s_%s", conf.Path, conf.Topic)
				_, ok := t.tskMap[mk]
				if ok {
					// 如果有，不需要操作
					continue
				} else {
					// 如果不在，启动一个taillog
					tailtask := NewTailTask(conf.Path, conf.Topic)
					// 将tailtask存入tskmap
					t.tskMap[mk] = tailtask
				}
				// 找出logEntry有，现在newConf没有的要删除,停止goroutine
				for _, c1 := range t.logEntry { // 原来的配置中拿出配置项
					isDelete := true
					for _, c2 := range newConf { // 从新的配置中拿出配置比较
						if c2.Path == c1.Path && c2.Topic == c1.Topic {
							isDelete = false
							continue
						}
					}
					if isDelete {
						// 把c1对应的tail删除
						mk := fmt.Sprintf("%s_%s", c1.Path, c1.Topic)
						t.tskMap[mk].cancelFunc() // 执行退出
					}
				}
			}
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
