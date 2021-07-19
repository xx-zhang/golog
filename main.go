package main

//import db "golog/models"
import (
	"fmt"
	conf "golog/handles"
	path "golog/utils"
	//"log"
	//"reflect"
	//"golog/tests"
)


func main() {
	var str1, str2 string
	str1 = path.GetCurrentDirectory()
	fmt.Println(str1)
	str2 = path.GetParentDirectory(str1)
	fmt.Println(str2)


	datas := conf.GetConf("D:\\home\\projects\\golog\\config.yaml")

	fmt.Println(datas.Mysql[0].Database)

}
