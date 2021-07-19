package logsource

import (
	"github.com/hpcloud/tail"
	"fmt"
	"time"
	conf "golog/handles"
)

func TailPath(path string)  {
	fmt.Printf(path)
}


func TailExample() {
	config := conf.GetConf("D:\\home\\projects\\golog\\config.yaml")
	logpath := config.Modsec[0].Logpath
	tailFile, err := tail.TailFile(logpath, tail.Config{
		ReOpen:    true,
		Follow:    true,
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll:      true,
	})

	if err != nil {
		fmt.Println("tail file err:", err)
		return
	}

	for true {
		msg, ok := <- tailFile.Lines
		if !ok {
			fmt.Printf("tail file close reopen, filename: %s\n", tailFile.Filename)
			time.Sleep(100 * time.Millisecond)
			continue
		}
		fmt.Println("msg:", msg)
	}
}