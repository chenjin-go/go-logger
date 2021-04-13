package logger

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
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
		fmt.Println("logger-conf.json not find ,use default config ,err:", readerr)
	}
	err := json.Unmarshal(by, &Config)
	if err != nil {
		fmt.Println("logger-conf.json content err:", err)
	}
	fmt.Println(Config)
	allfile, allerr := createDir(Config.AllLog)
	infofile, infoerr := createDir(Config.AllLog)
	errorfile, errorerr := createDir(Config.AllLog)
	warnfile, warnerr := createDir(Config.AllLog)

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
