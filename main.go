package main

import (
	"fmt"
	"github.com/hpcloud/tail"
	"golog/core"
	"golog/dao"
	conf "golog/handles"
	logging "golog/utils"
	"time"
)

var (
	tFlag   bool        // 时间是否在5s内重置过
	logpath string  // 日志存储路径
	maxnum int     // 单次插入日志的最大量
	timelen int    // 多久没有日志进来就会直接批量插入
	items []core.AuditLogItem
)

func init () {
	config := conf.GetConf("config.yaml")
	logpath = config.Modsec[0].LogPath
	maxnum = config.Modsec[0].MaxNum
	timelen = config.Modsec[0].TimeLen
	tFlag = false  // 默认5s内没有被重置，如果有数据流，但是不持续输出，那么定时器就会进行插入操作。

	// TODO 批量插入的函数- 一种是达到了时间，一种是达到了数量限制。
	// dao.BulkInsert(items)
	// items = [] core.AuditLogItem{}
}

func watchFile(msgChan chan<- string){
	// 生产者： 将观测到的 tail 行发送到消息管道
	fileWatcher, err := tail.TailFile(logpath, tail.Config{
		ReOpen:    true,
		Follow:    true,
		Location:  &tail.SeekInfo{Offset: 0, Whence: 1},
		MustExist: false,
		Poll:      false, // 这里如果改成true就读取不到了。
	})

	if err != nil {
		fmt.Println("tail file err:", err)
		return
	}

	for line := range fileWatcher.Lines {
		msgChan <- fmt.Sprintf(line.Text)
		logging.Info.Println("Sending Message to Msg Channel ---")
	}
}

func parseLine(msgChan <- chan string){
	for {
		message := <- msgChan
		amObj := core.ParseSingleLine(message)
		items = append(items, amObj)
		if len(items) > maxnum {
			dao.BulkInsert(items)
			items = [] core.AuditLogItem{}
			tFlag = true
		}
	}

}

func timerPoint() {
	// 时间针： 主要记录在这个时间段的心跳内
	heartbeat := time.Tick(time.Duration(timelen * 1000 * 1000 *1000))

	for {
		select {
		case <- heartbeat:
			// 没有插入过并且内存中的元素列表不为空
			if !tFlag {
				if len(items) > 0{
					dao.BulkInsert(items)
					items = [] core.AuditLogItem{}
				}
			}
			tFlag = false
			logging.Info.Println("程序 5s 心跳维持")
		}
	}
}

func main() {
	channel := make(chan string, 5) // 定义带有5个缓冲区的信道(当然可以是其他数字)
	go watchFile(channel) // 将 watchFile 函数交给协程处理, 产生的结果传入信道中
	go timerPoint()
	parseLine(channel) // 主线程从信道中取数据

}
