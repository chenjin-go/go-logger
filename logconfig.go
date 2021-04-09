package logger

type LogConfig struct {
	AllLog   string
	InforLog string
	ErrorLog string
	WarnLog  string
}

var Config = LogConfig{}

func init() {

}
