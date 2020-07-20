package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Configs struct{
	Master string `json:"master"`
	KeyName string `json:"keyName"`
	Port int `json:"port"`
	DbHost string `json:"dbHost"`
	DbUser string `json:"dbUser"`
	DbName string `json:"dbName"`
	DbPass string `json:"dbPass"`
}

type JsonStruct struct {
}

func NewJsonStruct() *JsonStruct {
	return &JsonStruct{}
}

func (c *JsonStruct) Load(fileName string, v interface{}) {
	data, err := ioutil.ReadFile(fileName)

	if err != nil {
		return
	}
	err = json.Unmarshal(data, v)
	if err != nil {
		return
	}
}

var Config Configs

func InitRedisConfigs(conf string) {

	fmt.Println("path",conf)

	JsonParse := NewJsonStruct()
	JsonParse.Load(conf, &Config)

}