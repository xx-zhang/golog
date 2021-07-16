package controllers

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Configs struct {
	Conf []struct {
		Name    string `yaml:"host"`
		Filter  string `yaml:"filter"`
		Pattern string `yaml:"pattern"`
		Rule    string `yaml:"rule"`
	} `yaml:"mysql"`
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
