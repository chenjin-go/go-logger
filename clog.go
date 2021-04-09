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

// const INFO = "INFO"
// const ERROR = "ERROR"
// const WARN = "WARN"

var ALLLOG = Config.AllLog
var INFOPATH = Config.InforLog
var ERRORPATH = Config.ErrorLog
var WARNPATH = Config.WarnLog

//创建日志文件
var ALLFile, _ = os.OpenFile(ALLLOG, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
var INFOFile, _ = os.OpenFile(INFOPATH, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
var ERRORFile, _ = os.OpenFile(ERRORPATH, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
var WARNFile, _ = os.OpenFile(WARNPATH, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)

var fileMap = map[Level]*os.File{
	INFO:  INFOFile,
	ERROR: ERRORFile,
	WARN:  WARNFile,
}

var typeMap = map[Level]string{
	INFO:  "INFO",
	ERROR: "ERROR",
	WARN:  "WARN",
}

type Level uint

type Clog struct {
	m     sync.Mutex
	buf   bytes.Buffer
	level Level
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
	_, err := l.buf.WriteString(header)
	if err != nil {
		fmt.Println(err)
	}
}

func (l *Clog) log(level Level, str string) {
	l.m.Lock()
	defer l.m.Unlock()
	l.level = level
	l.setHeader()
	l.buf.WriteString(str)
	l.save()
}

func init() {
	fmt.Println("测试")
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
	fileAndStdoutWriter.Write(l.buf.Bytes())
}
