package taillog

import (
	"fmt"

	"github.com/hpcloud/tail"
	"wklwukailun.com/logagent/kafka"
)

// 一个日志收集的任务
type TailTask struct {
	path     string
	topic    string
	instance *tail.Tail
}

func NewTailTask(path, topic string) (tailObj *TailTask) {
	tailObj = &TailTask{
		path:  path,
		topic: topic,
	}
	tailObj.init() //根据路径打开对应日志
	return
}
func (t *TailTask) init() {
	config := tail.Config{
		ReOpen:    true,                                 //重新打开
		Follow:    true,                                 //是否跟随
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, //从文件那个地方开始读
		MustExist: false,                                //文件不存在不报错
		Poll:      true,
	}
	var err error
	t.instance, err = tail.TailFile(t.path, config)
	if err != nil {
		fmt.Println("fail file failed, err:", err)
	}
	// 开启goroutine去法奥嗯kafka
	go t.run()
}

func (t *TailTask) run() {
	for {
		select {
		case line := <-t.instance.Lines: // 从通道中一行一行读取日志数据
			//kafka.SendToKafka(t.topic, line.Text) // 函数调用函数
			// 先将日志发送到通道中，在kafka包中有一个单独的goroutine去取数据并发送
			kafka.SendToChan(t.topic, line.Text)
		}
	}
}

// var (
// 	tailObj *tail.Tail
// )
// 专门从日志文件收集
// func Init(fileName string) (err error) {
// 	config := tail.Config{
// 		ReOpen:    true,                                 //重新打开
// 		Follow:    true,                                 //是否跟随
// 		Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, //从文件那个地方开始读
// 		MustExist: false,                                //文件不存在不报错
// 		Poll:      true,
// 	}
// 	tailObj, err = tail.TailFile(fileName, config)
// 	if err != nil {
// 		fmt.Println("fail file failed, err:", err)
// 		return
// 	}
// 	return
// }

// func ReadChan() <-chan *tail.Line {
// 	return tailObj.Lines
// }
