package logsource

import (
	"encoding/json"
	"fmt"
	"github.com/hpcloud/tail"
	conf "golog/handles"
)

type LineData struct {
	Datetime   string   `json:"datetime"`
	Uid string `json:"uid"`
	Index int `json:"index"`
}

// 创建一个 parser 的
func worker(start chan bool, index int) {
	<-start
	fmt.Println("This is Worker:", index)
}

func TailExample() {
	config := conf.GetConf("D:\\home\\projects\\golog\\config.yaml")
	logPath := config.Modsec[0].LogPath
	tailFile, err := tail.TailFile(logPath, tail.Config{
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

	for line := range tailFile.Lines {
		//
		//fmt.Println(line.Text)
		linda := &LineData{}
		err := json.Unmarshal([] byte(line.Text), linda)
		// 抛出的 panic 被自己 defer 语句中的 recover 捕获
		//fmt.Println(line.Text)
		fmt.Println(err)
		fmt.Println(*linda)
	}
	//fmt.Printf(msg.Text)

}