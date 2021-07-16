package main

import (
	"fmt"
	controllers "golog/handles"
	//yaml "gopkg.in/yaml.v2"
	jsoniter "github.com/json-iterator/go"
)

type AppInfo struct {
	Name string
}

func main() {
	info := AppInfo{
		Name: "GoApp",
	}
	jsonString, _ := jsoniter.Marshal(&info)

	fmt.Println(string(jsonString))
	controllers.ShowInfo()
}
