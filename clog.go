package logger

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	INFO = iota
	ERROR
	WARN
)

var ALLLOG = Config.AllLog
var INFOPATH = Config.InforLog
var ERRORPATH = Config.ErrorLog
var WARNPATH = Config.WarnLog

//创建日志文件
var ALLFile *os.File
var INFOFile *os.File
var ERRORFile *os.File
var WARNFile *os.File

var fileMap map[Level]*os.File

var typeMap = map[Level]string{
	INFO:  "INFO",
	ERROR: "ERROR",
	WARN:  "WARN",
}

type Level uint

func init() {

	allfile, _ := os.OpenFile(ALLLOG, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	infofile, _ := os.OpenFile(INFOPATH, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	errorfile, _ := os.OpenFile(ERRORPATH, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	warnfile, _ := os.OpenFile(WARNPATH, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)

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

type Clog struct {
	m        sync.Mutex
	Buf      bytes.Buffer
	level    Level
	TopHooks []itophook
	BotHooks []ibothook
}

func CreateLogger() *Clog {
	return &Clog{}
}

func (l *Clog) setHeader() {
	timestr := getTime()
	_, file, line, ok := runtime.Caller(3)
	if !ok {
		file = "???"
		line = 0
	}
	files := strings.Split(file, "/")
	header := timestr + " | " + typeMap[l.level] + " | " + files[len(files)-1] + ":" + strconv.Itoa(line) + " | "
	_, err := l.Buf.WriteString(header)
	if err != nil {
		fmt.Println(err)
	}
}

func (l *Clog) log(level Level, str string) {
	l.m.Lock()
	defer l.m.Unlock()
	l.level = level
	l.setHeader()
	l.Buf.WriteString(str)
	for _, v := range l.TopHooks {
		v.Call(l)
	}
	l.save()
	for _, v := range l.BotHooks {
		v.Call(l)
	}
	l.Buf.Reset()
}

func getTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func (l *Clog) Info(a ...interface{}) {
	l.log(INFO, fmt.Sprintln(a))
}

func (l *Clog) Error(a ...interface{}) {
	l.log(ERROR, fmt.Sprintln(a))
}

func (l *Clog) Warn(a ...interface{}) {
	l.log(WARN, fmt.Sprintln(a))
}

func (l *Clog) save() {
	writers := []io.Writer{
		os.Stdout,
		ALLFile,
		fileMap[l.level]}
	fileAndStdoutWriter := io.MultiWriter(writers...)
	// 创建新的log对象
	fileAndStdoutWriter.Write(l.Buf.Bytes())
}
