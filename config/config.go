package config

import (
	"encoding/json"
	"io/ioutil"
)

type DataSourse struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	UserName string `json:"username"`
	PassWord string `json:"password"`
	Database string `json:"database"`
	Options  string `json:"options"`
}

//数据库连接全局变量
var DataSourseConfig = new(DataSourse)

func ConfigInit(fileName string) (err error) {
	dataSourseBytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return
	}
	err = json.Unmarshal(dataSourseBytes, DataSourseConfig)
	if err != nil {
		return
	}
	return nil
}
