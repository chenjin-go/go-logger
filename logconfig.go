package logger

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
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
	configInit()
}

func configInit() {
	by, readerr := ioutil.ReadFile("conf/logger-conf.json")
	if readerr != nil {
		Logger.Warn("logger-conf.json not find ,use default config ,err:", readerr)
	}
	err := json.Unmarshal(by, &Config)
	if err != nil {
		Logger.Warn("logger-conf.json content err:", err)
	}
	allfile, allerr := createDir(Config.AllLog)
	infofile, infoerr := createDir(Config.InforLog)
	errorfile, errorerr := createDir(Config.ErrorLog)
	warnfile, warnerr := createDir(Config.WarnLog)

	if allerr != nil || infoerr != nil || warnerr != nil || errorerr != nil {
		Logger.Error("create log file err :", err)
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
		Logger.Error("create dir err:", direrr)
		return nil, direrr
	}
	filetype := strings.Split(file, ".")[1]
	if filetype != "log" {
		Logger.Error("create file type err:", filetype)
		return nil, errors.New("file type error")
	}
	startTimer(splitLog, dir)
	log, err := os.OpenFile(dir, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		Logger.Error("create file err:", err)
		return nil, err
	}
	return log, nil
}

func splitLog(dir string) {
	time := time.Now()
	oldfile, fileerr := os.OpenFile(dir, os.O_WRONLY, 0644)
	if fileerr != nil {
		Logger.Error("open old file err:", fileerr)
	}
	fi, err := oldfile.Stat()
	if err != nil {
		log.Println("stat fileinfo error")
	}
	oldTime := fi.ModTime()
	if time.YearDay() != oldTime.YearDay() {
		os.Rename(dir, dir+"."+getTime()+".log")
	}
}

//golang 定时器，启动的时候执行一次，以后每天晚上12点执行
func startTimer(f func(string), dir string) {
	go func() {
		for {
			f(dir)
			now := time.Now()
			// 计算下一个零点
			next := now.Add(time.Hour * 24)
			next = time.Date(next.Year(), next.Month(), next.Day(), 0, 0, 0, 0, next.Location())
			t := time.NewTimer(next.Sub(now))
			<-t.C
		}
	}()
}
