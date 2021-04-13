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

type Clog struct {
	m        sync.Mutex   //锁
	Buf      bytes.Buffer //日志流
	Level    Level        //级别
	topHooks []itophook   //头部钩子
	botHooks []ibothook   //尾部钩子
}

var clog = CreateLogger()

func CreateLogger() *Clog {
	return &Clog{}
}

//添加钩子
func (l *Clog) AddTopHook(hook itophook) {
	l.topHooks = append(l.topHooks, hook)
}

func (l *Clog) AddBotHook(hook ibothook) {
	l.botHooks = append(l.botHooks, hook)
}

//设置头部信息
func (l *Clog) setHeader() {
	timestr := getTime()
	_, file, line, ok := runtime.Caller(3)
	if !ok {
		file = "???"
		line = 0
	}
	files := strings.Split(file, "/")
	header := timestr + " | " + typeMap[l.Level] + " | " + files[len(files)-1] + ":" + strconv.Itoa(line) + " | "
	_, err := l.Buf.WriteString(header)
	if err != nil {
		fmt.Println(err)
	}
}

//log实现
func (l *Clog) log(level Level, str string) {
	l.m.Lock()
	defer l.m.Unlock()
	l.Level = level
	l.setHeader()
	l.Buf.WriteString(str)
	for _, v := range l.topHooks {
		v.TopCall(l)
	}
	l.save()
	for _, v := range l.botHooks {
		v.BotCall(l)
	}
	l.Buf.Reset()
}

//获取时间字符串
func getTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func (l *Clog) Info(a ...interface{}) {
	l.log(INFO, fmt.Sprintln(a...))
}

func (l *Clog) Error(a ...interface{}) {
	l.log(ERROR, fmt.Sprintln(a...))
}

func (l *Clog) Warn(a ...interface{}) {
	l.log(WARN, fmt.Sprintln(a...))
}

//发送流
func (l *Clog) save() {
	writers := []io.Writer{
		os.Stdout,
		ALLFile,
		fileMap[l.Level]}
	fileAndStdoutWriter := io.MultiWriter(writers...)
	fileAndStdoutWriter.Write(l.Buf.Bytes())
}
