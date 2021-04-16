# go-logger 日志管理

## 1.配置log配置文件(不配置也可以，使用默认配置，在本级目录下自动创建)
1. 创建文件夹conf
2. 创建json文件logger-conf.json
```
{
	"AllLog": "log/app.log",
	"InforLog": "log/app-info.log",
	"ErrorLog": "log/app-error.log",
	"WarnLog": "log/app-warn.log"
}
```

## 2.使用log打印

1. 建立全局参数 var log=CreateLogger()即可使用
2. 目前只支持三级
> INFO
> ERROR
> WARN


## 3. Hook（钩子）

1. 分为头部钩子和尾部钩子，分别继承TopHook或者BotHook
2. 使用方法，实现TopCall()或者BotCall()方法
3. 内部有Clog对象，其中Buf则是打印日志输入流
4. 添加钩子，在创建对象后，在init()方法中调用AddTopHook()或者AddBotHook()方法添加钩子

## 4. 分割日志
启动服务器会自动运行，启动会运行一次，以后每天0点运行一次，在路径下自动分隔当日和昨日日志分割
