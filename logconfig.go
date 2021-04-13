package logger

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"strings"
)

// log配置类
type LogConfig struct {
	AllLog   string `json:"AllLog"`
	InforLog string `json:"InforLog"`
	ErrorLog string `json:"ErrorLog"`
	WarnLog  string `json:"WarnLog"`
}

//默认配置
var Config = LogConfig{
	AllLog:   "app.log",
	InforLog: "app-info.log",
	ErrorLog: "app-error.log",
	WarnLog:  "app-warn.log",
}

//初始化，创建文件夹，创建log文件
func init() {
	by, readerr := ioutil.ReadFile("conf/logger-conf.json")
	if readerr != nil {
		clog.Warn("logger-conf.json not find ,use default config ,err:", readerr)
	}
	err := json.Unmarshal(by, &Config)
	if err != nil {
		clog.Warn("logger-conf.json content err:", err)
	}
	allfile, allerr := createDir(Config.AllLog)
	infofile, infoerr := createDir(Config.InforLog)
	errorfile, errorerr := createDir(Config.ErrorLog)
	warnfile, warnerr := createDir(Config.WarnLog)

	if allerr != nil || infoerr != nil || warnerr != nil || errorerr != nil {
		clog.Error("create log file err :", err)
	}

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

//创建文件夹方法
func createDir(dir string) (*os.File, error) {
	dirs := strings.Split(dir, "/")
	file := dirs[len(dirs)-1]
	dirnew := ""
	for i, dir := range dirs {
		if i != len(dirs)-1 {
			dirnew += dir + "/"
		}
	}
	direrr := os.MkdirAll(dirnew, os.ModePerm)
	if direrr != nil {
		clog.Error("create dir err:", direrr)
		return nil, direrr
	}

	filetype := strings.Split(file, ".")[1]
	if filetype != "log" {
		clog.Error("create file type err:", filetype)
		return nil, errors.New("file type error")
	}
	log, err := os.OpenFile(dir, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		clog.Error("create file err:", err)
		return nil, err
	}
	return log, nil
}
