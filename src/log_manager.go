package src

import (
	"fmt"
	"log"
	"runtime"
	"strconv"
	"time"
)

func NewLogManager() *LogManager {
	return &LogManager{
		list: make(map[string]*LogFile),
	}
}

// 日志文件队列管理器
type LogManager struct {
	list map[string]*LogFile
}

func (l *LogManager) Run() {
	log.Println("log manager starting")
	go func() {
		for {
			time.Sleep(5 * time.Second)
			fmt.Printf("当前日志监控协程数：%d\n", runtime.NumGoroutine()-2)
			l.check()
		}
	}()
	l.entry()
	for {
		time.Sleep(100000 * time.Second)
	}
}

// 添加
func (l *LogManager) entry() {
	// TODO 查找所有文件
	for i := 0; i < 2; i++ {
		LogFile := NewLogFile(strconv.Itoa(i))
		LogFile.Run()
		l.list[strconv.Itoa(i)] = LogFile
	}
}

// 检查日志
func (l *LogManager) check() {
	// TODO 日志文件是否有效
	// 日期是否有效，过期旧退出监控协程
	// 是否有在监控，没有就启动
	log.Println("checking log files")
	for k, v := range l.list {
		if l.checkDate(k) != true {
			v.Cancel()
			delete(l.list, k)

		}
	}
}

// 检查日期
func (l *LogManager) checkDate(path string) bool {
	if path == "1" {
		return false
	}
	return true
}
