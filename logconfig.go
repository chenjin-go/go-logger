package logger

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
