package logger

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type LogConfig struct {
	AllLog   string `json:"AllLog"`
	InforLog string `json:"InforLog"`
	ErrorLog string `json:"ErrorLog"`
	WarnLog  string `json:"WarnLog"`
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
	err := json.Unmarshal(by, &Config)
	if err != nil {
		fmt.Println("logger-conf.json解析错误,err:", err)
	}
	fmt.Println(Config)
	allfile, _ := os.OpenFile(Config.AllLog, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	infofile, _ := os.OpenFile(Config.InforLog, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	errorfile, _ := os.OpenFile(Config.ErrorLog, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	warnfile, _ := os.OpenFile(Config.WarnLog, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)

	ALLFile = allfile
	INFOFile = infofile
	ERRORFile = errorfile
	WARNFile = warnfile

	fileMap = map[Level]*os.File{
		INFO:  INFOFile,
		ERROR: ERRORFile,
		WARN:  WARNFile,
	}
}
