package utils

import (
	"io"
	"log"
	"os"
)

var (
	Info *log.Logger
	Warning *log.Logger
	Error * log.Logger
)

func init(){
	// https://www.cnblogs.com/oxspirt/p/11960517.html

	errFile,err:=os.OpenFile("golog_running.log",os.O_CREATE|os.O_WRONLY|os.O_APPEND,0666)
	if err!=nil{
		log.Fatalln("打开日志文件失败：",err)
	}
	log.SetPrefix("【WafAuditLog】")
	Info = log.New(os.Stdout,"Info:",log.Ldate | log.Ltime | log.Lshortfile)
	Warning = log.New(os.Stdout,"Warning:",log.Ldate | log.Ltime | log.Lshortfile)
	Error = log.New(io.MultiWriter(os.Stderr,errFile),"Error:",log.Ldate | log.Ltime | log.Lshortfile)

}
