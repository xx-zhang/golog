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
		LogPath    string `yaml:"logpath"`
	} `yaml:"modsec"`

}


func GetConf(filepath string) Configs {
	data, errData := ioutil.ReadFile(filepath)
	if errData != nil {
		fmt.Println(errData)
	}
	coifs := Configs{}
	errYaml := yaml.Unmarshal(data, &coifs)
	if errYaml != nil {
		fmt.Println(errYaml)
	}
	return coifs
}
