package main

import "mylogger"

func main() {
	log := mylogger.NewFileLog("Debug", "./", "test", 1024*10)
	name := "王森堡"
	id := 18
	for {
		log.Debug("傻屄id:%d,name:%s", id, name)
		log.Trace("傻屄id:%d,name:%s", id, name)
		log.Warning("傻屄id:%d,name:%s", id, name)
		log.Error("傻屄id:%d,name:%s", id, name)
		log.Fatal("傻屄id:%d,name:%s", id, name)
	}

}
