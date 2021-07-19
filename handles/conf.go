package controllers

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Configs struct {
	Mysql []struct {
		Host    string `yaml:"host"`
		Port  string `yaml:"port"`
		Database string `yaml:"database"`
		Username    string `yaml:"username"`
		Password    string `yaml:"password"`
	} `yaml:"mysql"`

	Modsec []struct {
		Logpath    string `yaml:"logpath"`
	} `yaml:"modsec"`

}


func GetConf(filepath string) Configs {
	data, err_data := ioutil.ReadFile(filepath)
	if err_data != nil {
		fmt.Println(err_data)
	}
	confs := Configs{}
	err_yaml := yaml.Unmarshal(data, &confs)
	if err_yaml != nil {
		fmt.Println(err_yaml)
	}
	return confs
}
