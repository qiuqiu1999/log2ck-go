package monitor

import (
	"fmt"
	"log"
	"path/filepath"
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
					go func () {
					for {
					m.parseConf()
					time.Sleep(5 * time.Second)
					72, 1-4        19%
					package monitor

					import (
					"bufio"
					"center-monitor/common/global"
					"fmt"
					"github.com/go-redis/redis"
					"io"
					"log"
					"os"
					"strconv"
					"strings"
					"syscall"
					"time"
					)

					type TailFile struct {
					fileName     string
					redisListKey string
					matchPrefix  string
					checkTime    time.Duration
					bufSize      int
					isMonitor    bool
					}

					func (this *TailFile) lock() {
					this.isMonitor = true
					}

					func (this *TailFile) unlock() {
					this.isMonitor = false
					}

					func (this *TailFile) isLocking() bool {
					return this.isMonitor
					}

					func (this *TailFile) getInode() (uint64, error) {
					fileInfo, err := os.Stat(this.fileName)
					if err != nil {
					return 0, err
					}
					stat := fileInfo.Sys().(*syscall.Stat_t)
					return stat.Ino, nil
					}

					func (this *TailFile) getKey() (string, error) {
					ino, err := this.getInode()
					if err != nil {
					return "", err
					}
					key := global.Ip + "_ino_" + strconv.Itoa(int(ino)) + "_" + this.fileName
					return key, nil
					"tail_file.go" 188L, 3756C                                                                                                                                      1, 1          顶端
					package monitor

					import (
					"bufio"
					"center-monitor/common/global"
					"fmt"
					"github.com/go-redis/redis"
					"io"
					"log"
					"os"
					"strconv"
					"strings"
					"syscall"
					"time"
					)

					type TailFile struct {
					fileName     string
					redisListKey string
					matchPrefix  string
					checkTime    time.Duration
					bufSize      int
					isMonitor    bool
					}

					func (this *TailFile) lock() {
					this.isMonitor = true
					}

					func (this *TailFile) unlock() {
					this.isMonitor = false
					}

					func (this *TailFile) isLocking() bool {
					return this.isMonitor
					}

					func (this *TailFile) getInode() (uint64, error) {
					fileInfo, err := os.Stat(this.fileName)
					if err != nil {
					return 0, err
					}
					stat := fileInfo.Sys().(*syscall.Stat_t)
					return stat.Ino, nil
					}

					func (this *TailFile) getKey() (string, error) {
					ino, err := this.getInode()
					if err != nil {
					return "", err
					}
					key := global.Ip + "_ino_" + strconv.Itoa(int(ino)) + "_" + this.fileName
					return key, nil
					"tail_file.go" 188L, 3756C                                                                                                                                      1, 1          顶端
					}
					line = line[:len(line)-drop]

					offset += int64(len(line) + 1 + drop)
					if err = this.saveOffset(rdb, offset); err != nil {
					log.Fatal(err)
					}

					if strings.Contains(line, this.matchPrefix) {
					marshal := strings.Replace(line, this.matchPrefix, "", 1)
					if _, err = rdb.LPush(this.redisListKey, marshal).Result(); err != nil {
					fmt.Printf("LPush err: %s", err)
					}
					}
					}
					this.unlock()
					return
					}()
					}
