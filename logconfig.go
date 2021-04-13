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
	by, err := ioutil.ReadFile("conf/logger-conf.json")
	if err != nil {
		fmt.Println("采用默认logger配置")
	}
	json.Unmarshal(by, Config)
}
