package controllers

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Configs struct {
	Mysql []struct {
		Host    string `yaml:"host"`
		Port  int `yaml:"port"`
		Database string `yaml:"database"`
		Username    string `yaml:"username"`
		Password    string `yaml:"password"`
	} `yaml:"mysql"`

	Modsec []struct {
		LogPath    string `yaml:"logpath"`
		MaxNum    int `yaml:"maxnum"`
		TimeLen    int `yaml:"timelen"`
	} `yaml:"modsec"`

}


func GetConf(filepath string) Configs {
	data, errData := ioutil.ReadFile(filepath)
	if errData != nil {
		fmt.Println(errData)
	}
	cfgs := Configs{}
	errYaml := yaml.Unmarshal(data, &cfgs)
	if errYaml != nil {
		fmt.Println(errYaml)
	}
	return cfgs
}
