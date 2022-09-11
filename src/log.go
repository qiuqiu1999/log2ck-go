package src

import (
	"context"
	"fmt"
	"log"
	"runtime"
	"time"
)

func NewLogFile(path string) *LogFile {
	ctx, cancel := context.WithCancel(context.Background())
	return &LogFile{
		Path:   path,
		Cancel: cancel,
		Ctx:    ctx,
	}
}

type LogFile struct {
	Path   string
	IsRun  bool
	Cancel context.CancelFunc
	Ctx    context.Context
}

func (l *LogFile) exit() {
	l.IsRun = false
}

func (l *LogFile) Run() {
	l.IsRun = true
	go func() {
		// 开始监控文件
		for {
			select {
			case <-l.Ctx.Done():
				fmt.Printf("%d 日志已过期\n", l.Path)
				runtime.Goexit()
			default:
				fmt.Printf("%s 读取新数据\n", l.Path)
				time.Sleep(2 * time.Second)
			}
		}
	}()

}

func (l *LogFile) monitor() {
	log.Println("在干了， 别催")
}
