package monitor

import (
	"center-monitor/common/global/setting"
	"fmt"
	"log"
	"path/filepath"
)

type Monitor struct {
	filesSetting *setting.FilesSetting
	fileList     map[string]*TailFile //监控的文件
}

func NewMonitor(setting *setting.FilesSetting) *Monitor {
	return &Monitor{
		filesSetting: setting,
		fileList:     make(map[string]*TailFile),
	}
}

func (m *Monitor) add(tailFile *TailFile) {
	m.fileList[tailFile.fileName] = tailFile
}

func (m *Monitor) isAdd(fileName string) bool {
	_, ok := m.fileList[fileName]
	if ok {
		return true
	}
	return false
}

// 解析路径
func (m *Monitor) parseConf() {
	for _, v := range m.filesSetting.List {
		paths, err := filepath.Glob(v.Path)
		if err != nil {
			log.Fatalf("monitor.parseConf err: %v", err)
		}

		for _, path := range paths {
			if m.isAdd(path) == false {
				fmt.Printf("add %s success \n", path)
				tf := &TailFile{
					fileName:     path,
					redisListKey: v.RedisListKey,
					matchPrefix:  v.MatchPrefix,
					"monitor.go" 150L, 3161C                                                                                                                                        1, 1          顶端
					package monitor

					import (
					"center-monitor/common/global/setting"
					"fmt"
					"github.com/go-redis/redis"
					"log"
					"os"
					"path/filepath"
					"time"
					)

					type Monitor struct {
					rdb          *redis.Client
					filesSetting *setting.FilesSetting
					fileList     map[string]*TailFile //监控的文件
					}

					func NewMonitor(rdb *redis.Client, setting *setting.FilesSetting) *Monitor {
					return &Monitor{
					rdb:          rdb,
					filesSetting: setting,
					fileList:     make(map[string]*TailFile),
					}
					}

					func (m *Monitor) add(tailFile *TailFile) {
					m.fileList[tailFile.fileName] = tailFile
					}

					func (m *Monitor) isAdd(fileName string) bool {
					_, ok := m.fileList[fileName]
					if ok {
					return true
					}
					return false
					}

					// 解析路径
					func (m *Monitor) parseConf() {
					for _, v := range m.filesSetting.List {
					paths, err := filepath.Glob(v.Path)
					if err != nil {
					log.Fatalf("monitor.parseConf err: %v", err)
					}

					for _, path := range paths {
					if m.isAdd(path) == false {
					fmt.Printf("add %s success \n", path)
					tf := &TailFile{
					fileName:     path,
					redisListKey: v.RedisListKey,
					matchPrefix:  v.MatchPrefix,
					checkTime:    v.CheckTime * time.Second,
					bufSize:      v.BufSize,
					isMonitor:    false,
					}
					//ino := tf.getInode()
					//offset := m.rdb.Get()
					m.add(tf)
					}
					}

					}
					}

					// 文件
					func (m *Monitor) checkFile(){
					go func (){
					for {
					m.parseConf()
					time.Sleep(5 * time.Second)
					72, 1-4        19%
					}

				// 监控文件
				/*func (m *Monitor) watcher() {
				    fmt.Println("Watcher start")
				    for _, v := range m.fileList {
				        go func(t *TailFile) {
				            var (
				                tails *tail.Tail
				                err   error
				                line  *tail.Line
				                ok    bool
				            )

				            // 正在监控
				            if t.isActivate() == true {
				                return
				            }

				            // 监控
				            if tails, err = tail.TailFile(t.FileName, tail.Config{
				                ReOpen:    true,                                 // 重新打开
				                Follow:    true,                                 // 是否跟随
				                Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, // 从文件的哪个地方开始读
				                MustExist: false,                                // 文件不存在不报错
				                Poll:      true,                                 // 监听新行，使用tail -f，这个参数非常重要
				            }); err != nil {
				                fmt.Println("tail file failed, err:", err)
				                return
				            } else {
				                t.activate()
				            }

				            for {
				                line, ok = <-tails.Lines
				                if !ok {
				                    fmt.Printf("tail file close reopen, filename:%s\n", tails.Filename)
				                    time.Sleep(time.Second)
				                    t.cancel()
				                    return
				                }

				                if strings.Contains(line.Text, t.MatchPrefix) {
				                    marshal := strings.Replace(line.Text, t.MatchPrefix, "", 1)
				                    //fmt.Println(marshal)
				                    if _, err = m.rdb.LPush(t.RedisListKey, marshal).Result(); err != nil {
				                        fmt.Printf("LPush err: %s", err)
				                    }
				                }
				            }
				        }(v)
				    }
				}*/
