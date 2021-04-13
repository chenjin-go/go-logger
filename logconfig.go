package logger

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type LogConfig struct {
	AllLog   string
	InforLog string
	ErrorLog string
	WarnLog  string
}

var Config = LogConfig{
	AllLog:   "app.log",
	InforLog: "app-info.log",
	ErrorLog: "app-error.log",
	WarnLog:  "app-warn.log",
}

func init() {
	by, readerr := ioutil.ReadFile("conf/logger-conf.json")
	if readerr != nil {
		fmt.Println("logger-conf.json未找到,采用默认logger配置,err:", readerr)
	}
	err := json.Unmarshal(by, Config)
	if err != nil {
		fmt.Println("logger-conf.json解析错误,err:", err)
	}
}
