package main

import (
	"fmt"
	"time"

	"github.com/hpcloud/tail"
)

// tailf的用法

func main() {
	fileName := "./my.log" //从该文件读取
	config := tail.Config{
		ReOpen:    true,                                 //重新打开
		Follow:    true,                                 //是否跟随
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, //从文件什么位置开始读取,记录位置
		MustExist: false,                                //日志文件不存在是否报错
		Poll:      true,                                 //
	}
	tails, err := tail.TailFile(fileName, config)
	if err != nil {
		fmt.Println("tail file failed, err:", err)
		return
	}
	var (
		line *tail.Line
		ok   bool
	)
	for {
		line, ok = <-tails.Lines //一行行读取
		if !ok {
			fmt.Println("tail file close reopen,filename:", tails.Filename)
			time.Sleep(time.Second)
			continue
		}
		fmt.Println("msg:", line.Text)
	}
}
