# go-logger 日志管理

## 1.配置log配置文件
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
